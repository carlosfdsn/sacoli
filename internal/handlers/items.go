package handlers

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)


type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Quantity int     `json:"quantity" binding:"required,gt=0"`
	Price    float64 `json:"price" binding:"gte=0"`
	Bought   bool    `json:"bought"`
}
var (
	mu     sync.Mutex
	items  = map[int]Item{}
	nextID = 1
)


// healthCheck godoc
// @Summary Verifica a saúde da API
// @Description Retorna o status de funcionamento da API.
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// listItems godoc
// @Summary Lista os itens
// @Description Retorna todos os itens cadastrados.
// @Tags items
// @Produce json
// @Success 200 {array} Item
// @Router /items [get]
func ListItems(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()
	c.JSON(http.StatusOK, items)
}

// getItemByID godoc
// @Summary Busca um item por ID
// @Description Retorna um único item a partir do seu ID.
// @Tags items
// @Produce json
// @Param id path int true "ID do item"
// @Success 200 {object} Item
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 404 {object} map[string]string "Item não encontrado"
// @Router /items/{id} [get]
func GetItemByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	item, ok := items[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
		return
	}

	c.JSON(http.StatusOK, item)
}

// createItem godoc
// @Summary Cria um item
// @Description Cadastra um novo item na lista. O ID é gerado automaticamente.
// @Tags items
// @Accept json
// @Produce json
// @Param item body Item true "Dados do item"
// @Success 201 {object} Item
// @Failure 400 {object} map[string]string "Corpo da requisição inválido"
// @Router /items [post]
func CreateItem(c *gin.Context) {
	var item Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	mu.Lock()
	defer mu.Unlock()
	item.ID = nextID
	nextID++

	items[item.ID] = item

	c.JSON(http.StatusCreated, item)
}

// deleteItem godoc
// @Summary Remove um item
// @Description Remove um item da lista a partir do seu ID.
// @Tags items
// @Param id path int true "ID do item"
// @Success 204
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 404 {object} map[string]string "Item não encontrado"
// @Router /items/{id} [delete]
func DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido",})
		return
	}

	mu.Lock()
	defer mu.Unlock()


	if _, ok := items[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
		return
	}
	delete(items, id)
	c.Status(http.StatusNoContent)

}

// updateItem godoc
// @Summary Atualiza um item
// @Description Substitui os dados de um item existente a partir do seu ID.
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "ID do item"
// @Param item body Item true "Novos dados do item"
// @Success 200 {object} Item
// @Failure 400 {object} map[string]string "ID ou corpo inválido"
// @Failure 404 {object} map[string]string "Item não encontrado"
// @Router /items/{id} [put]
func UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido",})
		return
	}

	var updated Item
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Corpo da requisição inválido"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := items[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
		return
	}

	updated.ID = id
	items[id] = updated

	c.JSON(http.StatusOK, updated)
}

// markAsBought godoc
// @Summary Marca um item como comprado
// @Description Define o campo bought do item como true.
// @Tags items
// @Produce json
// @Param id path int true "ID do item"
// @Success 200 {object} Item
// @Failure 400 {object} map[string]string "ID inválido"
// @Failure 404 {object} map[string]string "Item não encontrado"
// @Router /items/{id}/bought [patch]
func MarkAsBought(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido",})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	item, ok := items[id]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item não encontrado"})
		return
	}

	item.Bought = true
	items[id] = item

	c.JSON(http.StatusOK, item)
}