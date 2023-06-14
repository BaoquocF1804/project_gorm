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

const menuString string = "1. Press 1 to add a nhan vien\n2. Press 2 to delete a nhan vien\n3. Press 3 to add hoa don\n4. Press 4 to add san pham\n5.Press 0 to quit\n"

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
		case "3":
			addHoaDon(db)
			break
		case "4":
			addSanPham(db)
			break
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

func addNhanvien(db *gorm.DB) {
	s := model.Nhanvien{}

	fmt.Println("Ma NV: ")
	s.MANV, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.MANV = strings.TrimSuffix(s.MANV, "\n")

	fmt.Println("Ho va ten: ")
	s.HOTEN, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.HOTEN = strings.TrimSuffix(s.HOTEN, "\n")

	fmt.Println("SDT: ")
	s.SODT, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.SODT = strings.TrimSuffix(s.SODT, "\n")

	fmt.Println("Ngay vao lam: ")
	s.NGVL = time.Now()

	model.AddNhanvien(db, s)
	return
}

func deleteNV(db *gorm.DB) {
	var input string
	fmt.Println("Ma NV: ")
	fmt.Scanln(&input)
	model.DeleteNV(db, input)
}

func addHoaDon(db *gorm.DB) {
	s := model.Hoadon{}

	fmt.Println("MaKH: ")
	s.MAKH, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.MAKH = strings.TrimSuffix(s.MAKH, "\n")

	fmt.Println("MaNV: ")
	s.MANV, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.MANV = strings.TrimSuffix(s.MANV, "\n")

	fmt.Println("Tri gia: ")
	reader1 := bufio.NewReader(os.Stdin)
	var trigia float64
	fmt.Fscanf(reader1, "%f", &trigia)
	s.NGHD = time.Now()
	s.TRIGIA = int(trigia)

	fmt.Println(s)
	//  Retrieve the Nhanvien and Khachang from the database
	var nv model.Nhanvien
	if err := db.Where("MANV = ?", s.MAKH).Find(&nv).Error; err != nil {
		return
	}
	var kh model.Khachhang
	if err := db.Where("MAKH = ?", s.MAKH).Find(&kh).Error; err != nil {
		return
	}

	// Step 3: Append the new CreditCard to the User's CreditCards slice
	nv.Hoadon = append(nv.Hoadon, s)
	kh.Hoadon = append(kh.Hoadon, s)
	model.AddHoaDon(db, s)

	// Step 4: Cap nhat doanh so
	var doanhSo model.Khachhang
	err := db.Where("MAKH = ?", s.MAKH).Find(&doanhSo).Error
	if err != nil {
		return
	}
	doanhSoCu := doanhSo.DOANHSO
	doanhSoMoi := doanhSoCu + trigia
	db.Model(&model.Khachhang{}).Where("MAKH = ?", s.MAKH).Update("DOANHSO", doanhSoMoi)
}

// add a record to students table
func addSanPham(db *gorm.DB) {
	s := model.Sanpham{}

	fmt.Println("Ma MSP: ")
	s.MASP, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.MASP = strings.TrimSuffix(s.MASP, "\n")

	fmt.Println("Ten san pham: ")
	s.TENSP, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.TENSP = strings.TrimSuffix(s.TENSP, "\n")

	fmt.Println("DVT: ")
	s.DVT, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.DVT = strings.TrimSuffix(s.DVT, "\n")

	fmt.Println("Nuoc SX: ")
	s.NUOCSX, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	s.NUOCSX = strings.TrimSuffix(s.NUOCSX, "\n")

	fmt.Println("Gia: ")
	reader1 := bufio.NewReader(os.Stdin)
	var trigia float64
	fmt.Fscanf(reader1, "%f", &trigia)
	s.GIA = trigia

	model.AddSP(db, s)
	return
}
