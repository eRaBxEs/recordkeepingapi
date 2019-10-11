package controller

import (
	"log"
	"net/http"
	"recordkeeping/lib/model"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// Record ...
type Record struct {
	db  *pg.DB
	env *Environment
	log *zap.SugaredLogger
}

// used by the compiler to confirm if Record{} conforms to Handler
var _ Handler = &Record{}

// Init ... Here the Record is implementing the interface Handler
func (s *Record) Init(env *Environment, prefix string) error {
	s.env = env
	s.db = env.DB
	s.log = env.Log.Sugar()

	rtr := env.Rtr
	g := rtr.Group(prefix)
	p := g.Group("/record")

	// for portal
	p.POST("/expense", s.SaveExpense)
	p.POST("/income", s.SaveIncome)
	p.GET("/allexpenses", s.AllExpenses)
	p.GET("/allincomes", s.AllIncomes)
	p.GET("/balance", s.GetBalance)

	return nil
}

// SaveExpense ...
func (s *Record) SaveExpense(c echo.Context) error {

	expense := model.Expense{}

	if err := c.Bind(&expense); err != nil {
		return err
	}

	if err := expense.Save(s.db); err != nil {
		log.Printf("Error:%v", err)
		return err
	}

	return c.JSON(http.StatusOK, expense.ID)
}

// SaveIncome ...
func (s *Record) SaveIncome(c echo.Context) error {

	income := model.Income{}

	if err := c.Bind(&income); err != nil {
		return err
	}

	if err := income.Save(s.db); err != nil {
		log.Printf("Error:%v", err)
		return err
	}

	return c.JSON(http.StatusOK, income.ID)
}

// AllExpenses ...
func (s *Record) AllExpenses(c echo.Context) error {
	expense := model.Expense{}

	res, err := expense.GetAll(s.db)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// AllIncomes ...
func (s *Record) AllIncomes(c echo.Context) error {
	income := model.Income{}

	res, err := income.GetAll(s.db)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// GetBalance ...
func (s *Record) GetBalance(c echo.Context) error {

	income := model.Income{}
	expense := model.Expense{}

	var balance, sumExpenseAmount, sumIncomeAmount decimal.Decimal

	allIncomes, err := income.GetAll(s.db)
	if err != nil {
		return err
	}

	allExpenses, err := expense.GetAll(s.db)

	if err != nil {
		return err
	}

	for _, inc := range allIncomes {
		sumIncomeAmount = sumIncomeAmount.Add(inc.Amount)
	}

	for _, exp := range allExpenses {
		sumExpenseAmount = sumExpenseAmount.Add(exp.Amount)
	}

	balance = sumIncomeAmount.Sub(sumExpenseAmount)

	return c.JSON(http.StatusOK, balance)

}
