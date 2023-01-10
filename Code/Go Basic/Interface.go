package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId  int
    basicpay int
}

// Nhân viên permanent
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

// Nhân viên hợp đồng
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

/*
tổng chi phí được tính bằng cách duyệt qua từng phần tử của slice SalaryCalculator và tính tổng mức lương của từng nhân viên
*/
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {  
    pemp1 := Permanent{1, 5000, 20}
    pemp2 := Permanent{2, 6000, 30}
    cemp1 := Contract{3, 3000}
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)

}