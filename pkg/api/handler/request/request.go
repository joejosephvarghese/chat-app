package request

type UserSignUp struct {
	UserName  string `json:"user_name"  binding:"required,min=3,max=15"`
	FirstName string `json:"first_name"  binding:"required,min=2,max=50"`
	LastName  string `json:"last_name"  binding:"required,min=1,max=50"`
	Age       uint   `json:"age"  binding:"required,numeric,min=13"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password"  binding:"required,min=5,max=30"`
}
