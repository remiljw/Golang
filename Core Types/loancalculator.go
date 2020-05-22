package main

import (
	"fmt"
	"errors"
)

const (
	goodScore = 450
	lowScoreRatio =10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("Invali credit score")
	ErrIncome = errors.New("income invalid")
	ErrLoanAmount = errors.New("Loan Amount Invalid")
	ErrLoanTerm = errors.New("Loan term not a multiple of 12")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	interest := 20.0
	
	if creditScore >= goodScore {
		interest = 15.0
	}
	// check score
	if creditScore < 1 {
		return ErrCreditScore
	}
	// check income
	if income < 1 {
		return ErrIncome
	}
	// check loanamount
	if loanAmount < 1 {
		return ErrLoanAmount
	}
	// check term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) / loanTerm) + (loanAmount / loanTerm)
	// total cost of the loan
	totalInterest := (payment * loanTerm) - loanAmount
	approved := false
	if income > payment {
		// payment percent of income
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = true
		}
	}
	fmt.Println("Credit Score    :", creditScore)
	fmt.Println("Income          :", income)
	fmt.Println("Loan Amount     :", loanAmount)
	fmt.Println("Loan Term       :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate            :", interest)
	fmt.Println("Total Cost      :", totalInterest)
	fmt.Println("Approved        :", approved)
	fmt.Println("")

	return nil
}

func main() {
	// approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
// denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 12)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}