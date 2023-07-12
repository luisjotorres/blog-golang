package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math"
	"sync"
)

var wg sync.WaitGroup

func NewClient() (Client, error) {
	dns := "root:12345678@tcp(localhost:3306)/blog?parseTime=true"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &client{
		db,
	}, nil
}

func (c *client) Migrate(entity interface{}) {
	if err := c.AutoMigrate(entity); err != nil {
		panic(err)
	}
}

func (c *client) GetGORMClient() *gorm.DB {
	return c.DB
}

func (c *client) MigrateAndSeed(entity interface{}, payload ...interface{}) {
	var err error
	if err = c.AutoMigrate(entity); err == nil && c.Migrator().HasTable(entity) {
		if err := c.First(entity).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			wg.Add(len(payload))

			for _, item := range payload {
				c.seed(item)
			}
			wg.Wait()
		} else {
			fmt.Println("Estoy aca papu", errors.Is(err, gorm.ErrRecordNotFound))
		}
	} else {
		fmt.Println("Hubo un error :c", err, err != nil && c.Migrator().HasTable(entity))
	}
	fmt.Printf("Se han insertado %d registros \n", len(payload))
}

func (c *client) seed(payload interface{}) {
	defer wg.Done()
	c.Save(payload)
}

func (c *client) FindOneById(id uint, entity interface{}) interface{} {
	result := c.Where("id = ?", id).First(entity)
	return result
}

func (c *client) Save(entity interface{}) {
	err := c.Create(entity).Error
	if err != nil {
		fmt.Println(err)
	}
}

func (c *client) FindOneBy(key, value string, entity interface{}) {
	c.Where(fmt.Sprintf("%s = ?", key), value).First(entity)
}

func (c *client) Update(payload, entity interface{}) {
	c.Model(entity).Updates(payload)
}

func (c *client) FindAllByPage(entity interface{}, page int, pageNumber *int64) {
	//Retornaremos 10 por pagina
	c.Scopes(Paginate(page)).Find(entity)
	fmt.Println(entity)
	c.Model(entity).Count(pageNumber)
	fmt.Println(math.Ceil(float64(*pageNumber) / float64(10)))
	*pageNumber = int64(math.Ceil(float64(*pageNumber) / float64(10)))
}

func Paginate(pageNumber int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (pageNumber - 1) * 10
		return db.Offset(offset).Limit(10)
	}
}
