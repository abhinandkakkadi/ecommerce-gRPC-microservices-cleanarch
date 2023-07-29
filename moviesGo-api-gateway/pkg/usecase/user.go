package usecase

import (
	"errors"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	services "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/usecase/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type userUseCase struct {
	userRepo    interfaces.UserRepository
	cartRepo    interfaces.CartRepository
	productRepo interfaces.ProductRepository
}

func NewUserUseCase(repo interfaces.UserRepository, cartRepositiry interfaces.CartRepository, productRepository interfaces.ProductRepository) services.UserUseCase {
	return &userUseCase{
		userRepo:    repo,
		cartRepo:    cartRepositiry,
		productRepo: productRepository,
	}
}

func (u *userUseCase) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {

	// Check whether the user already exist. If yes, show the error message, since this is signUp
	userExist := u.userRepo.CheckUserAvailability(user.Email)

	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}

	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	// Hash password since details are validated
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	// add user details to the database
	userData, err := u.userRepo.UserSignUp(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// create referral code for the user and send in details of referred id of user if it exist
	id := uuid.New().ID()
	str := strconv.Itoa(int(id))
	userReferral := str[:8]
	err = u.userRepo.CreateReferralEntry(userData, userReferral)
	if err != nil {
		return models.TokenUsers{}, err
	}

	if user.ReferralCode != "" {
		// first check whether if a user with that referralCode exist
		referredUserId, err := u.userRepo.GetUserIdFromReferrals(user.ReferralCode)
		if err != nil {
			return models.TokenUsers{}, err
		}

		if referredUserId != 0 {
			referralAmount := 100
			err := u.userRepo.UpdateReferralAmount(float64(referralAmount), referredUserId, userData.Id)
			if err != nil {
				return models.TokenUsers{}, err
			}

		}
	}

	// crete a JWT token string for the user
	accessToken, err := helper.GenerateAccessToken(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	refreshToken, err := helper.GenerateRefreshToke(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	// copies all the details except the password of the user
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users:        userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *userUseCase) LoginHandler(user models.UserLogin) (models.TokenUsers, error) {

	// checking if a username exist with this email address
	ok := u.userRepo.CheckUserAvailability(user.Email)
	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist")
	}

	isBlocked, err := u.userRepo.UserBlockStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, err
	}

	if isBlocked {
		return models.TokenUsers{}, errors.New("user is not authorized to login")
	}

	// Get the user details in order to check the password, in this case ( The same function can be reused in future )
	user_details, err := u.userRepo.FindUserByEmail(user)
	if err != nil {
		return models.TokenUsers{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect")
	}

	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &user_details)
	if err != nil {
		return models.TokenUsers{}, err
	}

	// crete a JWT token string for the user
	accessToken, err := helper.GenerateAccessToken(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	refreshToken, err := helper.GenerateRefreshToke(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	return models.TokenUsers{
		Users:        userDetails,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}


// user checkout section
func (u *userUseCase) Checkout(userID int) (models.CheckoutDetails, error) {

	// list all address added by the user
	allUserAddress, err := u.userRepo.GetAllAddresses(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get available payment options
	paymentDetails, err := u.userRepo.GetAllPaymentOption()
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	walletDetails, err := u.userRepo.GetWalletDetails(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get all items from users cart
	cartItems, err := u.cartRepo.GetAllItemsFromCart(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}

	// get grand total of all the product
	grandTotal, err := u.cartRepo.GetTotalPrice(userID)
	if err != nil {
		return models.CheckoutDetails{}, err
	}


	return models.CheckoutDetails{
		AddressInfoResponse: allUserAddress,
		Payment_Method:      paymentDetails,
		Cart:                cartItems,
		Wallet:              walletDetails,
		Grand_Total:         grandTotal.TotalPrice,
		Total_Price:         grandTotal.FinalPrice,
	}, nil
}


func (u *userUseCase) AddAddress(address models.AddressInfo, userID int) error {

	err := u.userRepo.AddAddress(address, userID)
	if err != nil {
		return err
	}

	return nil
}



func (u *userUseCase) GetAllAddress(userID int) ([]models.AddressInfoResponse, error) {

	userAddress, err := u.userRepo.GetAllAddresses(userID)

	if err != nil {
		return []models.AddressInfoResponse{}, nil
	}

	return userAddress, nil

}



