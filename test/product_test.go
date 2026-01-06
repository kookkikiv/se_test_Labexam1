package test

import (
	"testing"
	"kookkikiv/se_test_Labexam1/model"
	"github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"

)

func TestBookingValidation(t *testing.T){
	g := NewgomegaWith(t){
		t.Run("Success Booking ", func(t *testing.T){
			booking :=model.Booking{
				CustomerName: "John Doe",
				RoomNumber:   205,
				GuestCount:   2,
				PhoneNumber:  "0123456789",
			}
			ok,err :=gocalidator.ValidateStruct(booking)
			g.Expect(ok).To(gomega.BeTrue())
			g.Expect(err).To(gomega.BeNil())
		})
	}
}
