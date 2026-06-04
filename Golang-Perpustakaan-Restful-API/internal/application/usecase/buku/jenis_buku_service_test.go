package buku_test

import (
	"testing"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/buku"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/buku/mocks"
	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var JenbukRepoInterfaceMock = &mocks.IBukuRepository{Mock: mock.Mock{}}
var JenbukServices = buku.BukuService{Repository: JenbukRepoInterfaceMock}

func TestInsertJenbuk(t *testing.T) {
	JenbukDummy := domain.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	JenbukRepoInterfaceMock.On("CreateJenisBuku", mock.AnythingOfType("domain.Jenis_Buku")).Return(JenbukDummy, nil).Once()
	JenbukInsert := domain.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	data, err := JenbukServices.CreateJenisBuku(JenbukInsert)
	assert.Equal(t, nil, err)
	assert.Equal(t, JenbukInsert, data)
}

func TestGetDataJenisBuku(t *testing.T) {
	t.Run("Get All data", func(t *testing.T) {
		JenbukDummy := []domain.Jenis_Buku{
			{
				IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
				JenisBuku: "ini jenis buku",
				Deskripsi: "ini deskripsi",
			},
		}
		JenbukRepoInterfaceMock.On("GetAllJenisBuku").Return(JenbukDummy, nil).Once()

		data, err := JenbukServices.GetAllJenisBuku()
		assert.Equal(t, nil, err)
		assert.Equal(t, JenbukDummy, data)
	})

	t.Run("Get Single data", func(t *testing.T) {
		JenbukDummy := domain.Jenis_Buku{
			IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
			JenisBuku: "ini jenis buku",
			Deskripsi: "ini deskripsi",
		}
		JenbukRepoInterfaceMock.On("GetJenisBukuById", JenbukDummy.IDJenis).Return(&JenbukDummy, nil).Once()

		data, err := JenbukServices.GetJenisBukuById(JenbukDummy.IDJenis)
		assert.Equal(t, nil, err)
		assert.Equal(t, &JenbukDummy, data)
	})

	t.Run("Get Single empty data", func(t *testing.T) {
		JenbukDummy := domain.Jenis_Buku{}
		JenbukRepoInterfaceMock.On("GetJenisBukuById", "01H019WHW4A45KH65AH9XW89PC").Return(&JenbukDummy, nil).Once()

		data, err := JenbukServices.GetJenisBukuById("01H019WHW4A45KH65AH9XW89PC")
		assert.Equal(t, nil, err)
		assert.Equal(t, &JenbukDummy, data)
	})

	t.Run("Get All find data", func(t *testing.T) {
		JenbukDummy := []domain.Jenis_Buku{
			{
				IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
				JenisBuku: "ini jenis buku",
				Deskripsi: "ini deskripsi",
			},
		}
		JenbukRepoInterfaceMock.On("CariJenisBuku", "ini jenis buku").Return(JenbukDummy, nil).Once()

		data, err := JenbukServices.FindJenisBuku("ini jenis buku")
		assert.Equal(t, nil, err)
		assert.Equal(t, JenbukDummy, data)
	})

	t.Run("Get All find data empty", func(t *testing.T) {
		JenbukDummy := []domain.Jenis_Buku{}
		JenbukRepoInterfaceMock.On("CariJenisBuku", "ini jenis buku").Return(JenbukDummy, nil).Once()

		data, err := JenbukServices.FindJenisBuku("ini jenis buku")
		assert.Equal(t, nil, err)
		assert.Equal(t, JenbukDummy, data)
	})
}

func TestUpdateDataJenbuk(t *testing.T) {
	var DummyEdit domain.Jenis_Buku
	JenbukRepoInterfaceMock.On("HitungDataJenisBuku", mock.Anything).Return(int64(1)).Once()
	JenbukRepoInterfaceMock.On("UpdateJenisBuku", mock.AnythingOfType("domain.Jenis_Buku")).Return(domain.Jenis_Buku{}, nil).Once()
	dataUpdate := domain.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	data, err := JenbukServices.UpdateJenisBuku(dataUpdate)
	assert.Equal(t, nil, err)
	assert.Equal(t, DummyEdit, data)
}

func TestHapusDataJenbuk(t *testing.T) {
	JenbukDummy := domain.Jenis_Buku{
		IDJenis:   "01H019WHW4A45KH65AH9XW89PC",
		JenisBuku: "ini jenis buku",
		Deskripsi: "ini deskripsi",
	}
	JenbukRepoInterfaceMock.On("HitungDataJenisBuku", mock.Anything).Return(int64(1)).Once()
	JenbukRepoInterfaceMock.On("DeleteJenisBuku", mock.Anything).Return(nil).Once()
	err := JenbukServices.HapusJenisBuku(JenbukDummy.IDJenis)
	assert.Equal(t, nil, err)
}

