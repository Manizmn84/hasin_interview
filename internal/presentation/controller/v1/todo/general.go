package todo

import (
	"net/http"

	"github.com/Manizmn84/hasin_interview/bootstrap"
	tododto "github.com/Manizmn84/hasin_interview/internal/application/dto/todo"
	"github.com/Manizmn84/hasin_interview/internal/application/usecase"
	"github.com/Manizmn84/hasin_interview/internal/presentation/controller/base"
	"github.com/gin-gonic/gin"
)

type TodoGeneralController struct {
	cfg         *bootstrap.Config
	todoService usecase.TodoService
}

func NewTodoGeneralController(cfg *bootstrap.Config, todoService usecase.TodoService) *TodoGeneralController {
	return &TodoGeneralController{
		cfg:         cfg,
		todoService: todoService,
	}
}

func (tgc *TodoGeneralController) Create(ctx *gin.Context) {
	type TodoCreateParam struct {
		Title string  `json:"title" validate:"required"`
		Dsc   string  `json:"dsc" validate:"required"`
		Np    float64 `json:"np" validate:"numeric"`
	}

	params := base.Validated[TodoCreateParam](ctx)

	createTodo := tododto.TodoCreateRequest{
		Np:    params.Np,
		Title: params.Title,
		Dsc:   params.Dsc,
	}

	res, err := tgc.todoService.CreateTodo(createTodo)

	if err != nil {
		panic(err)
	}

	trans := base.GetTranslator(ctx, tgc.cfg.Constant.Context.Translator)
	message, _ := trans.Translate("successMessage.createTodo")
	base.Response(ctx, http.StatusCreated, message, res)
}

func (tgc *TodoGeneralController) GetByID(ctx *gin.Context) {
	type TodoGetByIDParam struct {
		ID uint `uri:"id" validate:"required"`
	}

	params := base.Validated[TodoGetByIDParam](ctx)

	res, err := tgc.todoService.GetTodoByID(params.ID)
	if err != nil {
		panic(err)
	}

	trans := base.GetTranslator(ctx, tgc.cfg.Constant.Context.Translator)
	message, _ := trans.Translate("successMessage.getTodo")
	base.Response(ctx, http.StatusOK, message, res)
}

func (tgc *TodoGeneralController) List(ctx *gin.Context) {
	res, err := tgc.todoService.GetAllTodosSortedByNp()
	if err != nil {
		panic(err)
	}

	trans := base.GetTranslator(ctx, tgc.cfg.Constant.Context.Translator)
	message, _ := trans.Translate("successMessage.listTodo")
	base.Response(ctx, http.StatusOK, message, res)
}

func (tgc *TodoGeneralController) Update(ctx *gin.Context) {
	type TodoUpdateParam struct {
		ID     uint    `uri:"id" validate:"required"`
		Title  string  `json:"title" validate:"required"`
		Dsc    string  `json:"dsc" validate:"required"`
		Np     float64 `json:"np" validate:"numeric"`
		Status uint    `json:"status" validate:"required,oneof=1 2 3"`
	}

	params := base.Validated[TodoUpdateParam](ctx)

	updateTodo := tododto.TodoUpdateRequest{
		ID:     params.ID,
		Np:     params.Np,
		Title:  params.Title,
		Dsc:    params.Dsc,
		Status: params.Status,
	}

	err := tgc.todoService.UpdateTodo(updateTodo)
	if err != nil {
		panic(err)
	}

	trans := base.GetTranslator(ctx, tgc.cfg.Constant.Context.Translator)
	message, _ := trans.Translate("successMessage.updateTodo")
	base.Response(ctx, http.StatusOK, message, nil)
}

func (tgc *TodoGeneralController) Delete(ctx *gin.Context) {
	type TodoDeleteByID struct {
		ID uint `uri:"id" validate:"required"`
	}

	params := base.Validated[TodoDeleteByID](ctx)

	err := tgc.todoService.DeleteTodo(params.ID)
	if err != nil {
		panic(err)
	}

	trans := base.GetTranslator(ctx, tgc.cfg.Constant.Context.Translator)
	message, _ := trans.Translate("successMessage.deleteTodo")
	base.Response(ctx, http.StatusOK, message, nil)
}
