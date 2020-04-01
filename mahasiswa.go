package tugasakhir

import (
	"database/sql"
	"project/belajargolang/tugasakhir/lib"
)

type Mahasiswa struct {
	NIM   string `json:"nim"`
	Nama  string `json:"nama"`
	Kelas string `json:"kelas"`
}

var TblMahasiswa = `CREATE TABLE mahasiswa (
  nim VARCHAR(20) PRIMARY KEY,
  nama VARCHAR(20),
  kelas VARCHAR(20)
);`

func (m *Mahasiswa) Name() string {
	return "mahasiswa"
}

func (m *Mahasiswa) Fields() (fields []string, dst []interface{}) {
	fields = []string{"nim", "nama", "kelas"}
	dst = []interface{}{&m.NIM, &m.Nama, &m.Kelas}
	return fields, dst
}

func (m *Mahasiswa) PrimaryKey() (fields []string, dst []interface{}) {
	fields = []string{"nim"}
	dst = []interface{}{&m.NIM}
	return fields, dst
}

func (m *Mahasiswa) Structur() lib.Table {
	return &Mahasiswa{}
}

func (m *Mahasiswa) Insert(db *sql.DB) error {
	return lib.Insert(db, m)
}

func (m *Mahasiswa) Update(db *sql.DB, change map[string]interface{}) (map[string]interface{}, error) {
	return change, lib.Update(db, m, change)
}

func (m *Mahasiswa) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}

func (m *Mahasiswa) Get(db *sql.DB) error {
	return lib.Get(db, m)
}

func Mahasiswas(db *sql.DB, params lib.RequestParams) ([]*Mahasiswa, error) {
	m := &Mahasiswa{}
	res, err := lib.Fetch(db, m, params)
	if err != nil {
		return nil, err
	}
	mhs := make([]*Mahasiswa, len(res))
	for index, item := range res {
		mhs[index] = item.(*Mahasiswa)
	}
	return mhs, nil
}
