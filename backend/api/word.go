package api

import (
	"net/http"
	"strconv"

	"ai-wordbook/model"
	"ai-wordbook/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryWordRequest struct {
	Word       string `json:"word" binding:"required"`
	AIProvider string `json:"ai_provider" binding:"required"`
}

type SaveWordRequest struct {
	Word       string                   `json:"word" binding:"required"`
	Definition string                   `json:"definition" binding:"required"`
	AIProvider string                   `json:"ai_provider" binding:"required"`
	Sentences  []service.SentenceResult `json:"sentences" binding:"required"`
}

func QueryWord(c *gin.Context) {
	var req QueryWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数校验失败: " + err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	var word model.Word
	err := model.DB.Where("user_id = ? AND word = ?", userID, req.Word).
		Preload("Sentences").First(&word).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"saved":       true,
			"word":        word.Word,
			"definition":  word.Definition,
			"ai_provider": word.AIProvider,
			"sentences":   word.Sentences,
		})
		return
	}

	provider, err := service.GetProvider(req.AIProvider)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := provider.QueryWord(req.Word)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "AI 查询失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"saved":       false,
		"word":        req.Word,
		"definition":  result.Definition,
		"ai_provider": req.AIProvider,
		"sentences":   result.Sentences,
	})
}

func SaveWord(c *gin.Context) {
	var req SaveWordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数校验失败: " + err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	var existing model.Word
	if err := model.DB.Where("user_id = ? AND word = ?", userID, req.Word).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "该单词已保存"})
		return
	}

	word := model.Word{
		UserID:     userID,
		Word:       req.Word,
		Definition: req.Definition,
		AIProvider: req.AIProvider,
	}
	for _, s := range req.Sentences {
		word.Sentences = append(word.Sentences, model.Sentence{
			English: s.English,
			Chinese: s.Chinese,
		})
	}

	if err := model.DB.Create(&word).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "保存成功", "id": word.ID})
}

func ListWords(c *gin.Context) {
	userID := c.GetUint("user_id")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var total int64
	model.DB.Model(&model.Word{}).Where("user_id = ?", userID).Count(&total)

	var words []model.Word
	model.DB.Where("user_id = ?", userID).
		Preload("Sentences").
		Order("created_at DESC").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&words)

	c.JSON(http.StatusOK, gin.H{
		"total":     total,
		"page":      page,
		"page_size": pageSize,
		"words":     words,
	})
}

func DeleteWord(c *gin.Context) {
	userID := c.GetUint("user_id")
	wordID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的单词 ID"})
		return
	}

	var word model.Word
	if err := model.DB.Where("id = ? AND user_id = ?", wordID, userID).First(&word).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "单词不存在或无权操作"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败"})
		return
	}

	model.DB.Where("word_id = ?", wordID).Delete(&model.Sentence{})
	model.DB.Delete(&word)

	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
