package repository

type Repository interface {
    // Define methods for data access
    FindByID(id string) (interface{}, error)
    Save(entity interface{}) error
    Delete(id string) error
    // Add more methods as needed
}