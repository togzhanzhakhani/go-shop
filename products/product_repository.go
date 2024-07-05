package products

import(
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func(pr *ProductRepository) SaveProduct(product *Product) error {
	return pr.DB.Create(&product).Error
}

func (pr *ProductRepository) GetAllProducts() ([]Product, error) {
	var products []Product
	err := pr.DB.Find(&products).Error
	return products, err
}

func (pr *ProductRepository) GetProductByID(id string) (*Product, error) {
	var product Product
	err := pr.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (pr *ProductRepository) UpdateProduct(id string, updatedProduct *Product) error {
	return pr.DB.Model(&Product{}).Where("id = ?", id).Updates(updatedProduct).Error
}

func (pr *ProductRepository) DeleteProduct(id string) error {
	return pr.DB.Delete(&Product{}, id).Error
}

func (pr *ProductRepository) SearchProductsByName(name string) ([]Product, error) {
	var products []Product
	err := pr.DB.Where("name ILIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

func (pr *ProductRepository) SearchProductsByCategory(category string) ([]Product, error) {
	var products []Product
	err := pr.DB.Where("category ILIKE ?", "%"+category+"%").Find(&products).Error
	return products, err
}