package controller

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"os"
	"project_demo/cmd/model"
	"time"
)

const menuString string = "1. Press 1 to add a nhan vien\n2. Press 2 to delete a nhan vien\n3. Press 3 to add hoa don\n4.Press 0 to quit\n"

func Menu(db *gorm.DB) {
	for true {
		var inp string
		fmt.Print(menuString)
		fmt.Scanln(&inp)
		if inp == "0" {
			break
		}
		switch inp {
		case "1":
			addNhanvien(db)
			break
		//case "2":
		//	addStudent(db)
		//	break
		case "3":
			addHoaDon(db)
			break
		//case "4":
		//	lendBook(db)
		//	break
		//case "5":
		//	returnBook(db)
		//	break
		//case "6":
		//	getBookByName(db)
		//	break
		//case "7":
		//	getBookByID(db)
		//	break
		//case "8":
		//	listBooksStudentKeep(db)
		//	break
		case "2":
			deleteNV(db)
		default:
			break
		}
		fmt.Println("\nEnter 1 to continue\nEnter other character to quit")
		fmt.Scanln(&inp)
		if inp == "1" {
			continue
		} else {
			break
		}
	}
}

// add a record to students table
func addNhanvien(db *gorm.DB) {
	s := model.Nhanvien{}
	fmt.Println("Ma NV: ")
	s.MANV, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Ho va ten: ")
	s.HOTEN, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("SDT: ")
	s.SODT, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Ngay vao lam: ")
	s.NGVL = time.Now()
	fmt.Println(s)
	model.AddNhanvien(db, s)
	return
}

func deleteNV(db *gorm.DB) {
	var input string

	fmt.Println("Ma NV: ")

	fmt.Scanln(&input)
	model.DeleteNV(db, input)
}

// add a record to hoa don table
func addHoaDon(db *gorm.DB) {
	// Step 1: Create a new Hoadon instance
	fmt.Println("SoHD: ")
	reader := bufio.NewReader(os.Stdin)
	var num int
	_, err := fmt.Fscanf(reader, "%d", &num)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	fmt.Println("MaKH: ")
	makh, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("MaNV: ")
	manv, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println("Tri gia: ")
	reader1 := bufio.NewReader(os.Stdin)
	var trigia int
	_, err = fmt.Fscanf(reader1, "%d", &trigia)
	if err != nil {
		fmt.Println("Error reading integer:", err)
		return
	}

	s := model.Hoadon{
		SOHD:   num,
		NGHD:   time.Now(),
		MAKH:   makh,
		MANV:   manv,
		TRIGIA: trigia,
	}
	fmt.Println(s)

	// Step 2: Retrieve the Nhanvien from the database
	var nv model.Nhanvien
	if err := db.Where("member_number = ?", manv).First(&nv).Error; err != nil {
		return
	}

	// Step 3: Append the new CreditCard to the User's CreditCards slice
	nv.Hoadon = append(nv.Hoadon, s)
	model.AddHoaDon(db, s)
	return
}
