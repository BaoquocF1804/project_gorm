package controller

import (
	"bufio"
	"fmt"
	"gorm.io/gorm"
	"os"
	"project_demo/cmd/model"
	"strings"
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
	fmt.Println("MaKH: ")
	makh, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	makh = strings.TrimSuffix(makh, "\n")
	fmt.Println("MaNV: ")
	manv, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	manv = strings.TrimSuffix(manv, "\n")
	fmt.Println("Tri gia: ")
	reader1 := bufio.NewReader(os.Stdin)
	var trigia float64
	fmt.Fscanf(reader1, "%f", &trigia)

	//trigia := float64(60000)
	s := model.Hoadon{
		NGHD:   time.Now(),
		MAKH:   makh,
		MANV:   manv,
		TRIGIA: int(trigia),
	}
	fmt.Println(s)

	// Step 2: Retrieve the Nhanvien and Khachang from the database

	var nv model.Nhanvien
	if err := db.Where("MANV = ?", manv).Find(&nv).Error; err != nil {
		return
	}

	var kh model.Khachhang
	if err := db.Where("MAKH = ?", makh).Find(&kh).Error; err != nil {
		return
	}
	// Step 3: Append the new CreditCard to the User's CreditCards slice
	nv.Hoadon = append(nv.Hoadon, s)
	kh.Hoadon = append(kh.Hoadon, s)

	model.AddHoaDon(db, s)

	// Step 3: Cap nhat doanh so
	var doanhSo model.Khachhang
	err := db.Where("MAKH = ?", makh).Find(&doanhSo).Error
	if err != nil {
		return
	}
	doanhSoCu := doanhSo.DOANHSO
	doanhSoMoi := doanhSoCu + trigia
	db.Model(&model.Khachhang{}).Where("MAKH = ?", makh).Update("DOANHSO", doanhSoMoi)
}
