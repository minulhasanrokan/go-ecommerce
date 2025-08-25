package services

import (
	"errors"

	"github.com/minulhasanrokan/go-ecommerce/cmd/api/requests"
	"github.com/minulhasanrokan/go-ecommerce/cmd/internal/models"
)

func (service *AppService) FindProductById(id uint) (*models.ProductModel, error) {

	var product *models.ProductModel
	err := service.Db.First(&product, id).Error

	if err != nil {

		return nil, errors.New("product does not exist")
	}

	return product, nil
}

func (service *AppService) GetCategories() ([]*models.CategoryModel, error) {

	categories, err := service.FindCategories()

	if err != nil {
		return nil, errors.New("categories does not exist")
	}

	return categories, err
}

func (service *AppService) FindCategories() ([]*models.CategoryModel, error) {

	var categories []*models.CategoryModel

	err := service.Db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (service *AppService) GetCategory(id uint) (*models.CategoryModel, error) {

	cat, err := service.FindCategoryById(id)

	if err != nil {
		return nil, errors.New("category does not exist")

	}

	return cat, nil
}

func (service *AppService) FindCategoryById(id uint) (*models.CategoryModel, error) {

	var category *models.CategoryModel

	err := service.Db.First(&category, id).Error

	if err != nil {

		return nil, errors.New("category does not exist")
	}

	return category, nil
}

func (service *AppService) CreateCategory(payload *requests.CreateCategoryRequest) (models.CategoryModel, error) {

	category, err := service.CreateCategoryData(&models.CategoryModel{
		Name:         payload.Name,
		ImageUrl:     payload.ImageUrl,
		DisplayOrder: payload.DisplayOrder,
	})

	if err != nil {

		return models.CategoryModel{}, err
	}

	return category, nil
}

func (service *AppService) CreateCategoryData(model *models.CategoryModel) (models.CategoryModel, error) {

	err := service.Db.Create(&model).Error

	if err != nil {
		return models.CategoryModel{}, errors.New("create category failed")
	}

	return *model, nil
}

func (service *AppService) EditCategory(id uint, input *requests.CreateCategoryRequest) (*models.CategoryModel, error) {

	exitCat, err := service.FindCategoryById(id)

	if err != nil {
		return nil, errors.New("category does not exist")

	}

	if len(input.Name) > 0 {
		exitCat.Name = input.Name
	}

	if input.ParentId > 0 {
		exitCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		exitCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		exitCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := service.UpdateCategory(exitCat)

	return updatedCat, err
}

func (service *AppService) UpdateCategory(e *models.CategoryModel) (*models.CategoryModel, error) {

	err := service.Db.Save(&e).Error

	if err != nil {
		return nil, errors.New("fail to update category")
	}

	return e, nil
}

func (service *AppService) DeleteCategory(id uint) error {

	err := service.Db.Delete(&models.CategoryModel{}, id).Error

	if err != nil {
		return errors.New("fail to delete category")
	}

	return nil
}

func (service *AppService) CreateProduct(input *requests.CreateProductRequest, user models.UserModel) (models.ProductModel, error) {
	product, err := service.CreateNewProduct(&models.ProductModel{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		ImageUrl:    input.ImageUrl,
		UserId:      int(user.Id),
		Stock:       uint(input.Stock),
	})

	if err != nil {
		return models.ProductModel{}, err
	}

	return product, nil
}

func (service *AppService) CreateNewProduct(e *models.ProductModel) (models.ProductModel, error) {

	err := service.Db.Model(&models.ProductModel{}).Create(e).Error

	if err != nil {

		return models.ProductModel{}, errors.New("cannot create product")
	}

	return *e, err
}

func (service *AppService) GetProducts() ([]*models.ProductModel, error) {

	products, err := service.FindProducts()

	if err != nil {

		return nil, errors.New("products does not exist")
	}

	return products, err
}

func (service *AppService) FindProducts() ([]*models.ProductModel, error) {

	var products []*models.ProductModel

	err := service.Db.Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (service *AppService) GetProductById(id uint) (*models.ProductModel, error) {

	product, err := service.FindProductById(id)

	if err != nil {

		return nil, errors.New("product does not exist")
	}

	return product, nil
}

func (service *AppService) EditProduct(id uint, input *requests.CreateProductRequest, user models.UserModel) (*models.ProductModel, error) {

	exitProduct, err := service.FindProductById(id)

	if err != nil {
		return nil, errors.New("product does not exist")
	}

	// verify product owner
	if exitProduct.UserId != int(user.Id) {

		return nil, errors.New("you don't have manage rights of this product")
	}

	if len(input.Name) > 0 {
		exitProduct.Name = input.Name
	}

	if len(input.Description) > 0 {
		exitProduct.Description = input.Description
	}

	if input.Price > 0 {
		exitProduct.Price = input.Price
	}

	if input.CategoryId > 0 {
		exitProduct.CategoryId = input.CategoryId
	}

	updatedProduct, err := service.UpdateProduct(exitProduct)

	return updatedProduct, err
}

func (service *AppService) UpdateProduct(e *models.ProductModel) (*models.ProductModel, error) {

	err := service.Db.Save(&e).Error

	if err != nil {

		return nil, errors.New("fail to update product")
	}
	return e, nil
}

func (service *AppService) UpdateProductStock(e models.ProductModel, productId uint) (*models.ProductModel, error) {

	product, err := service.FindProductById(productId)

	if err != nil {
		return nil, errors.New("product not found")
	}

	if product.UserId != e.UserId {

		return nil, errors.New("you don't have manage rights of this product")
	}

	product.Stock = e.Stock

	editProduct, err := service.UpdateProduct(product)

	if err != nil {
		return nil, err
	}

	return editProduct, nil
}

func (service *AppService) DeleteProduct(id uint, user models.UserModel) error {

	exitProduct, err := service.FindProductById(id)

	if err != nil {
		return errors.New("product does not exist")
	}

	if exitProduct.UserId != int(user.Id) {
		return errors.New("you don't have manage rights of this product")
	}

	err = service.DeleteProductData(exitProduct)

	if err != nil {
		
		return errors.New("product cannot delete")
	}

	return nil
}

func (service *AppService) DeleteProductData(e *models.ProductModel) error {

	err := service.Db.Delete(&models.ProductModel{}, e.Id).Error

	if err != nil {
		return errors.New("product cannot delete")
	}

	return nil
}
