package unit
import(
	"testing"
	"github.com/se/examle/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run("Field Ea is required",func(t *testing.T){

		sl := entity.E{
			Ea: "",
		}

		ok, err := govalidator.ValidateStruct(sl)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(ContainSubstring("Field Ea is required"))
	})
}