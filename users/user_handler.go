package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"shop/validation"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	UserRepo *UserRepository
}

func NewUserHandler(ur *UserRepository) *UserHandler {
	return &UserHandler{UserRepo: ur}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
        return
    }

    if err := validation.ValidateStruct(&user); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), UserBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

    if _, err := uh.UserRepo.getUserByEmail(user.Email); err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
        return
    }

    if err := uh.UserRepo.saveUser(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully!"})
}

func (uh *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := uh.UserRepo.getAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving users"})
		return
	}

	if len(users) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
        return
    }

	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.UserRepo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updatedUser User
	if err := c.BindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error decoding request body"})
		return
	}

	if err := validation.ValidateStruct(&updatedUser); err != nil {
		errorMessage := validation.HandleValidationErrors(err.(validator.ValidationErrors), UserBaseMessages)
		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
		return
	}

	existingUser, err := uh.UserRepo.GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if updatedUser.Email != existingUser.Email {
        if _, err := uh.UserRepo.getUserByEmail(updatedUser.Email); err == nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
            return
        }
    }

	if err := uh.UserRepo.updateUser(id, &updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully!"})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := uh.UserRepo.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := uh.UserRepo.deleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}

func (uh *UserHandler) SearchUsersByName(c *gin.Context) {
	name := c.Query("name")
	users, err := uh.UserRepo.searchUsersByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching users by name"})
		return
	}

	if len(users) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
        return
    }

	c.JSON(http.StatusOK, users)
}

func (uh *UserHandler) SearchUsersByEmail(c *gin.Context) {
	email := c.Query("email")
	users, err := uh.UserRepo.searchUsersByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching users by email"})
		return
	}

	if len(users) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
        return
    }

	c.JSON(http.StatusOK, users)
}
