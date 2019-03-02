package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

//tipe data utk data train, penggunaan json belum pasti
type Train struct {
	Id           string `json:"id"`
	Age          string `json:"age"`
	Workclass    string `json:"workclass"`
	Education    string `json:"education"`
	Marital      string `json:"marital"`
	Occupation   string `json:"occupation"`
	Relationship string `json:"relation"`
	Hour         string `json:"hour"`
	Income       string `json:"income"`
}

func main() {
	//import csv
	csvFile, _ := os.Open("TrainsetTugas1ML.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var trainml []Train
	i := 0
	for {

		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if i == 0 {
			i++
			continue
		}
		trainml = append(trainml, Train{
			Id:           line[0],
			Age:          line[1],
			Workclass:    line[2],
			Education:    line[3],
			Marital:      line[4],
			Occupation:   line[5],
			Relationship: line[6],
			Hour:         line[7],
			Income:       line[8],
		})
	}
	fmt.Println("data train telah di import")
	//trainJson, _ := json.Marshal(trainml)
	//fmt.Println(string(trainJson))
	//declarasi var
	//ditambah 1 untuk smothing lapalace
	// var age & class
	var fYoungR float32 = 1
	var fAdultR float32 = 1
	var fOldR float32 = 1

	var fYoungP float32 = 1
	var fAdultP float32 = 1
	var fOldP float32 = 1

	// var workclass & class
	var fPrivateR float32 = 1
	var fGovR float32 = 1
	var fSelfR float32 = 1

	var fPrivateP float32 = 1
	var fGovP float32 = 1
	var fSelfP float32 = 1

	// var education& class
	var fSomeR float32 = 1 //some colage
	var fGradR float32 = 1 //HSgraduate
	var fBachR float32 = 1 //Bachelors

	var fSomeP float32 = 1 //some colage
	var fGradP float32 = 1 //HSgraduate
	var fBachP float32 = 1 //Bachelors

	//var marital-status & class
	var fMarrR float32 = 1 //Maried
	var fNevrR float32 = 1 //Never married
	var fDivdR float32 = 1 //Divorced

	var fMarrP float32 = 1 //Maried
	var fNevrP float32 = 1 //Never married
	var fDivdP float32 = 1 //Divorced

	//var occupation & class
	var fProfR float32 = 1 //Prof-speciality
	var fCrftR float32 = 1 //Craft-repair
	var fExecR float32 = 1 //Exec-managerial

	var fProfP float32 = 1
	var fCrftP float32 = 1
	var fExecP float32 = 1

	//var relationship
	var fHusbR float32 = 1 // Husband
	var fNfamR float32 = 1 //Not-in-Family
	var fOchdR float32 = 1 // Own-Child

	var fHusbP float32 = 1
	var fNfamP float32 = 1
	var fOchdP float32 = 1

	//var Hours-per-week
	var fNormR float32 = 1 //Normal
	var fManyR float32 = 1 //Many
	var fLowR float32 = 1  //Low

	var fNormP float32 = 1
	var fManyP float32 = 1
	var fLowP float32 = 1

	//Mulai penghitungan frequency
	for j := 0; j < 160; j++ {
		//freq age & class
		if trainml[j].Income == ">50K" {
			if trainml[j].Age == "young" {
				fYoungR++
			}
			if trainml[j].Age == "adult" {
				fAdultR++
			}
			if trainml[j].Age == "old" {
				fOldR++
			}
		} else {
			if trainml[j].Age == "young" {
				fYoungP++
			}
			if trainml[j].Age == "adult" {
				fAdultP++
			}
			if trainml[j].Age == "old" {
				fOldP++
			}
		}

		//freq workclass & class
		if trainml[j].Income == ">50K" {
			if trainml[j].Workclass == "Private" {
				fPrivateR++
			}
			if trainml[j].Workclass == "Self-emp-not-inc" {
				fSelfR++
			}
			if trainml[j].Workclass == "Local-gov" {
				fGovR++
			}
		} else {
			if trainml[j].Workclass == "Private" {
				fPrivateP++
			}
			if trainml[j].Workclass == "Self-emp-not-inc" {
				fSelfP++
			}
			if trainml[j].Workclass == "Local-gov" {
				fGovP++
			}
		}

		//freq education & class
		if trainml[j].Income == ">50K" {
			if trainml[j].Education == "Some-college" {
				fSomeR++
			}
			if trainml[j].Education == "HS-grad" {
				fGradR++
			}
			if trainml[j].Education == "Bachelors" {
				fBachR++
			}
		} else {
			if trainml[j].Education == "Some-college" {
				fSomeP++
			}
			if trainml[j].Education == "HS-grad" {
				fGradP++
			}
			if trainml[j].Education == "Bachelors" {
				fBachP++
			}
		}

		//freq marital-status & class
		if trainml[j].Income == ">50K" {
			if trainml[j].Marital == "Married-civ-spouse" {
				fMarrR++
			}
			if trainml[j].Marital == "Never-married" {
				fNevrR++
			}
			if trainml[j].Marital == "Divorced" {
				fDivdR++
			}
		} else {
			if trainml[j].Marital == "Married-civ-spouse" {
				fMarrP++
			}
			if trainml[j].Marital == "Never-married" {
				fNevrP++
			}
			if trainml[j].Marital == "Divorced" {
				fDivdP++
			}
		}

		//freq occupation &class
		if trainml[j].Income == ">50K" {
			if trainml[j].Occupation == "Prof-specialty" {
				fProfR++
			}
			if trainml[j].Occupation == "Craft-repair" {
				fCrftR++
			}
			if trainml[j].Occupation == "Exec-managerial" {
				fExecR++
			}
		} else {
			if trainml[j].Occupation == "Prof-specialty" {
				fProfP++
			}
			if trainml[j].Occupation == "Craft-repair" {
				fCrftP++
			}
			if trainml[j].Occupation == "Exec-managerial" {
				fExecP++
			}
		}

		//freq relationship &class
		if trainml[j].Income == ">50K" {
			if trainml[j].Relationship == "Husband" {
				fHusbR++
			}
			if trainml[j].Relationship == "Not-in-family" {
				fNfamR++
			}
			if trainml[j].Relationship == "Own-child" {
				fOchdR++
			}
		} else {
			if trainml[j].Relationship == "Husband" {
				fHusbP++
			}
			if trainml[j].Relationship == "Not-in-family" {
				fNfamP++
			}
			if trainml[j].Relationship == "Own-child" {
				fOchdP++
			}
		}

		//freq hours &class
		if trainml[j].Income == ">50K" {
			if trainml[j].Hour == "normal" {
				fNormR++
			}
			if trainml[j].Hour == "many" {
				fManyR++
			}
			if trainml[j].Hour == "low" {
				fLowR++
			}
		} else {
			if trainml[j].Hour == "normal" {
				fNormP++
			}
			if trainml[j].Hour == "many" {
				fManyP++
			}
			if trainml[j].Hour == "low" {
				fLowP++
			}
		}

	}
	fmt.Println("frequency data sudah di hitung")

	//probalitas dengan laplace
	//P Age >50K
	var pAdultR = fAdultR / (fAdultR + fOldR + fYoungR)
	var pOldR = fOldR / (fAdultR + fOldR + fYoungR)
	var pYoungR = fYoungR / (fAdultR + fOldR + fYoungR)
	//P AGe <=50K
	var pAdultP = fAdultP / (fAdultP + fOldP + fYoungP)
	var pOldP = fOldP / (fAdultP + fOldP + fYoungP)
	var pYoungP = fYoungP / (fAdultP + fOldP + fYoungP)

	//P workclass >50K
	var pPrivateR = fPrivateR / (fPrivateR + fSelfR + fGovR)
	var pSelfR = fSelfR / (fPrivateR + fSelfR + fGovR)
	var pGovR = fGovR / (fPrivateR + fSelfR + fGovR)
	//P workclass <=50K
	var pPrivateP = fPrivateP / (fPrivateP + fSelfP + fGovP)
	var pSelfP = fSelfP / (fPrivateP + fSelfP + fGovP)
	var pGovP = fGovP / (fPrivateP + fSelfP + fGovP)

	//P Education >50K
	var pSomeR = fSomeR / (fSomeR + fGradR + fBachR)
	var pGradR = fGradR / (fSomeR + fGradR + fBachR)
	var pBachR = fBachR / (fSomeR + fGradR + fBachR)
	//P education <=50K
	var pSomeP = fSomeP / (fSomeP + fGradP + fBachP)
	var pGradP = fGradP / (fSomeP + fGradP + fBachP)
	var pBachP = fBachP / (fSomeP + fGradP + fBachP)

	//P Marital >50K
	var pMarrR = fMarrR / (fMarrR + fNevrR + fDivdR)
	var pNevrR = fNevrR / (fMarrR + fNevrR + fDivdR)
	var pDivdR = fDivdR / (fMarrR + fNevrR + fDivdR)
	//P Marital <=50K
	var pMarrP = fMarrP / (fMarrP + fNevrP + fDivdP)
	var pNevrp = fNevrP / (fMarrP + fNevrP + fDivdP)
	var pDivdP = fDivdP / (fMarrP + fNevrP + fDivdP)

	//P Occupation >50K
	var pProfR = fProfR / (fProfR + fCrftR + fExecR)
	var pCrftR = fCrftR / (fProfR + fCrftR + fExecR)
	var pExecR = fExecR / (fProfR + fCrftR + fExecR)
	//P Occupation <=50K
	var pProfP = fProfP / (fProfP + fCrftP + fExecP)
	var pCrftP = fCrftP / (fProfP + fCrftP + fExecP)
	var pExecP = fExecP / (fProfP + fCrftP + fExecP)

	//P Relationship >50K
	var pHusbR = fHusbR / (fHusbR + fNfamR + fOchdR)
	var pNfamR = fNfamR / (fHusbR + fNfamR + fOchdR)
	var pOchdR = fOchdR / (fHusbR + fNfamR + fOchdR)
	//P Relationship <=50K
	var pHusbP = fHusbP / (fHusbP + fNfamP + fOchdP)
	var pNfamP = fNfamP / (fHusbP + fNfamP + fOchdP)
	var pOchdP = fOchdP / (fHusbP + fNfamP + fOchdP)

	//P Hours >50K
	var pNormR = fNormR / (fNormR + fManyR + fLowR)
	var pManyR = fManyR / (fNormR + fManyR + fLowR)
	var pLowR = fLowR / (fNormR + fManyR + fLowR)
	//P hours <=50K
	var pNormP = fNormP / (fNormP + fManyP + fLowP)
	var pManyP = fManyP / (fNormP + fManyP + fLowP)
	var pLowP = fLowP / (fNormP + fManyP + fLowP)
	fmt.Println("nilai probabilitas data sudah di hitung")

	//import TestsetTugas1ML.csv
	csvFile, _ = os.Open("TestsetTugas1ML.csv")
	reader = csv.NewReader(bufio.NewReader(csvFile))
	var testml []Train
	i = 0
	for {

		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if i == 0 {
			i++
			continue
		}
		testml = append(testml, Train{
			Id:           line[0],
			Age:          line[1],
			Workclass:    line[2],
			Education:    line[3],
			Marital:      line[4],
			Occupation:   line[5],
			Relationship: line[6],
			Hour:         line[7],
		})
	}
	fmt.Println("data test sudah di import")
	//data di tangkap di var testml
	var InR float32 //utk >50K
	var InP float32 //utk <=50K
	for j := 0; j < 40; j++ {
		InR = 1
		InP = 1
		//cek Age
		if testml[j].Age == "young" {
			InR = InR * pYoungR
			InP = InR * pYoungP
		} else if testml[j].Age == "adult" {
			InR = InR * pAdultR
			InP = InR * pAdultP
		} else {
			InR = InR * pOldR
			InP = InR * pOldP
		}

		//Cek Workclass
		if testml[j].Workclass == "Private" {
			InR = InR * pPrivateR
			InP = InP * pPrivateP
		} else if testml[j].Workclass == "Self-emp-not-inc" {
			InR = InR * pSelfR
			InP = InP * pSelfP
		} else {
			InR = InR * pGovR
			InP = InP * pGovP
		}

		//cek education
		if testml[j].Workclass == "HS-grad" {
			InR = InR * pGradR
			InP = InP * pGradP
		} else if testml[j].Workclass == "Bachelors" {
			InR = InR * pBachR
			InP = InP * pBachP
		} else {
			InR = InR * pSomeR
			InP = InP * pSomeP
		}

		//cek marital
		if testml[j].Marital == "Married-civ-spouse" {
			InR = InR * pMarrR
			InP = InP * pMarrP
		} else if testml[j].Marital == "Divorced" {
			InR = InR * pDivdR
			InP = InP * pDivdP
		} else {
			InR = InR * pNevrR
			InP = InP * pNevrp
		}

		//cek occupation
		if testml[j].Occupation == "Prof-specialty" {
			InR = InR * pProfR
			InP = InP * pProfP
		} else if testml[j].Occupation == "Craft-repair" {
			InR = InR * pCrftR
			InP = InP * pCrftP
		} else {
			InR = InR * pExecR
			InP = InP * pExecP
		}

		//cek relationship
		if testml[j].Relationship == "Husband" {
			InR = InR * pHusbR
			InP = InP * pHusbP
		} else if testml[j].Relationship == "Not-in-family" {
			InR = InR * pNfamR
			InP = InP * pNfamP
		} else {
			InR = InR * pOchdR
			InP = InP * pOchdP
		}

		//cek Hour
		if testml[j].Hour == "normal" {
			InR = InR * pNormR
			InP = InP * pNormP
		} else if testml[j].Hour == "many" {
			InR = InR * pManyR
			InP = InP * pManyP
		} else {
			InR = InR * pLowR
			InP = InP * pLowP
		}

		//kesimpulan income
		if InR > InP {
			testml[j].Income = ">50K"
		} else {
			testml[j].Income = "<=50K"
		}
	}
	fmt.Println("data train sudah memiliki nilai income")
	// testJson, _ := json.Marshal(trainml)
	// Data test sudah di uji, tinggal di export ke TebakanTugas1ML.csv
	file, err := os.Create("TebakanTugas1ML.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	// for _, value := range testml {
	// 	err := writer.Write(string(value))
	// 	checkError("Cannot write to file", err)
	// }
	var records = [][]string{
		{"id", "age", "workclass", "education", "marital-status", "occupation", "relationship", "hours-per-week", "Income"},
	}
	err = writer.WriteAll(records)
	for j := 0; j < 40; j++ {
		records = [][]string{
			{testml[j].Id, testml[j].Age, testml[j].Workclass, testml[j].Education, testml[j].Marital, testml[j].Occupation, testml[j].Relationship, testml[j].Hour, testml[j].Income},
		}
		err = writer.WriteAll(records)
		checkError("Cannot write to file", err)
	}
	fmt.Println("data sudah di export. Eksekusi selesai")
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
