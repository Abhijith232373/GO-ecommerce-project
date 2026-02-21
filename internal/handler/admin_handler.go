package handler

import (
	"e-commerce/internal/models"
	"e-commerce/internal/service"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminService service.AdminService
	ProductService service.ProductService
}

func NewAdminHandler(adminService *service.AdminService ,productService service.ProductService)*AdminHandler{
	return &AdminHandler{
		AdminService: *adminService,
		ProductService: productService,
	}
}

func (h *AdminHandler)ShowLoginPage(c *gin.Context){
	c.HTML(http.StatusOK,"login.html",nil)
}
func (h *AdminHandler)Login(c *gin.Context){
	email:=c.PostForm("email")
	password :=c.PostForm("password")

	admin,err := h.AdminService.Login(email,password)

	if err !=nil{
		c.HTML(http.StatusUnauthorized,"login.html",gin.H{
			"error":"Invalid email or password",
		})
		return
	}

	c.SetCookie(
		"admin_id",
		fmt.Sprintf("%d",admin.ID),
		3600,
		"/",
		"",
		false,
		true,
	)
	c.Redirect(http.StatusSeeOther,"/admin/dashboard")
}

func (h *AdminHandler)Logout(c *gin.Context){
	c.SetCookie(
		"admin_id",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
	c.Redirect(http.StatusSeeOther,"/admin/login")
}

func (h *AdminHandler)ShowDashboard(c *gin.Context){
	c.HTML(http.StatusOK,"dashboard.html",gin.H{
		"title":"Admin Dashboard",
	})
}


//////////////// Product management ///////////////////
func (h *AdminHandler)ShowProductsPage(c *gin.Context){
	search :=c.Query("search")
	var products []models.Product
	var err error
	// products,err:=h.AdminService.GetAllProducts()
	if search != ""{
		products,err =h.AdminService.SearchProducts(search)
	}else {
		products,err =h.AdminService.GetAllProducts()
	}
	if err !=nil{
		c.String(500,"Failed ")
		return
	}
	c.HTML(200,"products.html",gin.H{
		"Products":products,
		"Search":search,
	})
}

func (h *AdminHandler)AddProduuct(c *gin.Context){
	name:=c.PostForm("name")
	description :=c.PostForm("description")
	priceStr:=c.PostForm("price")
	stockStr:=c.PostForm("stock")

	price,_:=strconv.ParseFloat(priceStr,64)
	stock,_:=strconv.Atoi(stockStr)
	categoryIDStr :=c.PostForm("category_id")
	categoryID,_:=strconv.Atoi(categoryIDStr)

	file,err:=c.FormFile("image")
	if err !=nil{
		c.String(400,"Image upload failed")
		return
	}

	filename:=filepath.Base(file.Filename)
	savePath :="uploads/products/"+filename

	err =c.SaveUploadedFile(file,savePath)
	if err !=nil{
		c.String(500,"Save Failed")
		return
	}
	err =h.AdminService.AddProduct(
		name,
		description,
		price,
		stock,
		uint(categoryID),
		"products/"+filename,
	)

	// c.SaveUploadedFile(file,"upload/products/"+filename)

	// h.AdminService.AddProduct(name,price,stock,filename)

	if err !=nil{
		c.String(500,"DB save failed")
		return
	}
	c.Redirect(303,"/admin/products")
}

func (h *AdminHandler)DeleteProduct(c *gin.Context){
	id:=c.Param("id")

	h.AdminService.DeleteProduct(id)
	
	c.Redirect(303,"/admin/products")
}

func  (h *AdminHandler)ActivateProduct(c *gin.Context){
	id :=c.Param("id")

	err:=h.ProductService.ActivateProduct(id)
	if err !=nil{
		c.String(500,"Failed to  Activate")
		return
	}
	c.Redirect(302,"/admin/products")
}

func  (h *AdminHandler)SuspendProduct(c *gin.Context){
	id :=c.Param("id")
	err :=h.ProductService.SuspendProduct(id)

	if err !=nil{
		c.String(500,"Failed to susupend")
		return
	}
	c.Redirect(302,"/admin/products")
}


///////////////////// User management //////////////////////////////
func (h *AdminHandler)ShowUserPage(c *gin.Context){
	users,err:=h.AdminService.GetAllUsers()

	if err != nil {
		c.String(500,"failed to fetch  users")
		return
	}

		fmt.Println("DEBUG USERS:", users) 
	c.HTML(200,"users.html",gin.H{
		"Users":users,
	})
}

func (h *AdminHandler)GetUsers(c *gin.Context){
	users,err:=h.AdminService.GetAllUsers()
	if err !=nil{
		c.String(500,"Error fetching users")
		return
	}
	c.HTML(200,"admin/users.html",gin.H{
		"Users":users,
	})
}

func (h *AdminHandler)AddUser(c *gin.Context){

	name:=c.PostForm("name")
	email:=c.PostForm("email")
	password:=c.PostForm("password")

	err :=h.AdminService.AddUser(name,email,password)

	if err !=nil{
		c.JSON(500,gin.H{"error":err.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther,"/admin/users")
}

func  (h *AdminHandler)SuspendUser(c *gin.Context){
	id :=c.Param("id")

	h.AdminService.SuspendUser(id)

	c.Redirect(http.StatusSeeOther,"/admin/users")
}

func (h *AdminHandler)ActiveUser(c *gin.Context){
	id :=c.Param("id")

	h.AdminService.ActiveUser(id)

	c.Redirect(http.StatusSeeOther,"/admin/users")
}

func (h *AdminHandler)DeleteUser(c *gin.Context){
	id :=c.Param("id")

	if id == "" {
		c.String(400,"User Id is required")
		return
	}
	err :=h.AdminService.DeleteUser(id)

	if err!=nil{
		c.String(500,"Failed to delete user")
		return
	}
	c.Redirect(302,"/admin/users")
}

