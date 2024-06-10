package helper

import (
    "errors"
    "strconv"
)

// ValidateName melakukan validasi terhadap input name
// ValidateName melakukan validasi terhadap input name


// ValidateName melakukan validasi terhadap input name
func ValidateName(name string) error {
    if name == "" {
        return errors.New("	name tidak boleh kosong")
    }
    if _, err := strconv.Atoi(name); err == nil {
        return errors.New("name harus berupa string")
    }
    return nil
}



// ValidateLocation melakukan validasi terhadap input location
func ValidateLocation(location string) error {
    if location == "" {
        return errors.New("location tidak boleh kosong")
    }

	if _, err := strconv.Atoi(location); err == nil {
        return errors.New("Location harus berupa string")
    }
  
    return nil
}

// ValidateCategory melakukan validasi terhadap input category
func ValidateCategory(category string) error {
    if category == "" {
        return errors.New("category tidak boleh kosong")
    }
    
    // Memeriksa apakah kategori bisa diubah menjadi angka
    _, err := strconv.Atoi(category)
    if err == nil {
        return errors.New("category harus berupa string")
    }
    
    return nil
}

// ValidateQty melakukan validasi terhadap input qty
func ValidateQty(qty int) error {
    if qty <= 0 { // Memeriksa apakah qty kosong atau negatif
        return errors.New("qty tidak boleh kosong dan harus lebih dari 0")
    }
    // Validasi lainnya sesuai kebutuhan
    return nil
}

// ValidatePrice melakukan validasi terhadap input price
func ValidatePrice(price float64) error {
    if price <= 0 { // Memeriksa apakah price kosong atau negatif
        return errors.New("price tidak boleh kosong dan harus lebih dari 0")
    }
    // Validasi lainnya sesuai kebutuhan
    return nil
}
