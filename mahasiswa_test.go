package tugasakhir

import (
	"database/sql"
	"fmt"
	"project/belajargolang/tugasakhir/lib"
	"testing"
)

var user, password, host, port, dbname string

func init() {
	user = "root"
	password = ""
	host = "127.0.0.1"
	port = "3306"
	dbname = "sistemkampus"
}

func TestMahasiswa(t *testing.T) {
	db, err := InitDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	data := []*Mahasiswa{
		&Mahasiswa{NIM: "1711444", Nama: "Filza", Kelas: "2KA36"},
		&Mahasiswa{NIM: "17114445", Nama: "Filza", Kelas: "2KA36"},
		&Mahasiswa{NIM: "17114446", Nama: "Filza", Kelas: "2KA36"},
	}

	t.Run("Test Insert And Get", func(t *testing.T) {
		for _, item := range data {
			if err := item.Insert(db); err != nil {
				t.Fatal(err)
			}

			got := &Mahasiswa{NIM: item.NIM}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareMahasiswa(t, got, item)
		}

	})

	t.Run("Test Update And Get", func(t *testing.T) {
		update := map[string]interface{}{
			"nama": "AIRTAS",
		}

		dataUpdate := data[0]
		_, err := dataUpdate.Update(db, update)
		if err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Gets", func(t *testing.T) {
		result, err := Mahasiswas(db, lib.RequestParams{})
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := &Mahasiswa{NIM: item.NIM}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareMahasiswa(t, got, item)
		}
	})

	t.Run("Testing Delete", func(t *testing.T) {
		m := &Mahasiswa{NIM: data[0].NIM}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
		fmt.Println(m)
	})
}

func InitDb() (*sql.DB, error) {
	db, err := lib.Connect(user, password, host, port, "")
	if err != nil {
		return nil, err
	}
	if err := lib.DropDB(db, dbname); err != nil {
		return nil, err
	}

	if err := lib.CreateDatabase(db, dbname); err != nil {
		return nil, err
	}

	db, err = lib.Connect(user, password, host, port, dbname)
	if err != nil {
		return nil, err
	}

	if err = lib.Use(db, dbname); err != nil {
		return nil, err
	}

	if err = lib.CreateTable(db, TblMahasiswa); err != nil {
		return nil, err
	}

	return db, err
}

func compareMahasiswa(t *testing.T, got, want *Mahasiswa) {
	if got.NIM != want.NIM {
		t.Fatalf("got : %s want :%s npm tidak sama", got.NIM, want.NIM)
	}
	if got.Nama != want.Nama {
		t.Fatalf("got :%s want :%s Nama tidak Sama", got.Nama, want.Nama)
	}
	if got.Kelas != want.Kelas {
		t.Fatalf("got :%s want :%s Nama tidak Sama", got.Kelas, want.Kelas)
	}

}
