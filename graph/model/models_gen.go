// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddAll struct {
	Product    *Product    `json:"Product"`
	Company    *Company    `json:"Company"`
	Category   *Category   `json:"Category"`
	Style      *Style      `json:"Style"`
	Type       *Type       `json:"Type"`
	Department *Department `json:"Department"`
	Reviews    *Reviews    `json:"Reviews"`
	BuyFrom    *BuyFrom    `json:"BuyFrom"`
	Image      *Image      `json:"Image"`
}

type AddAllInput struct {
	Department *Department `json:"Department"`
}

type Address struct {
	AddressLine1 *string   `json:"Address_Line1"`
	AddressLine2 *string   `json:"Address_Line2"`
	AddressTo    *string   `json:"Address_To"`
	City         *string   `json:"City"`
	Zip          *string   `json:"Zip"`
	State        *USStates `json:"State"`
	Country      *string   `json:"Country"`
}

type AddressInput struct {
	AddressLine1 *string   `json:"Address_Line1"`
	AddressLine2 *string   `json:"Address_Line2"`
	AddressTo    *string   `json:"Address_To"`
	City         *string   `json:"City"`
	Zip          *string   `json:"Zip"`
	State        *USStates `json:"State"`
	Country      *string   `json:"Country"`
	ID           *string   `json:"ID"`
}

type BuyFrom struct {
	ID        *string  `json:"ID"`
	Link      *string  `json:"Link"`
	Price     *float64 `json:"Price"`
	LogoID    *string  `json:"Logo_ID"`
	CompanyID *string  `json:"Company_ID"`
}

type BuyFromInput struct {
	ID        *string  `json:"ID"`
	Link      *string  `json:"Link"`
	Price     *float64 `json:"Price"`
	LogoID    *string  `json:"Logo_ID"`
	CompanyID *string  `json:"Company_ID"`
}

type Category struct {
	ID           *string `json:"ID"`
	Title        *string `json:"Title"`
	DepartmentID *string `json:"Department_ID"`
	StyleID      *string `json:"Style_ID"`
	TypeID       *string `json:"Type_ID"`
}

type CategoryInput struct {
	ID           *string `json:"ID"`
	Title        *string `json:"Title"`
	DepartmentID *string `json:"Department_ID"`
	StyleID      *string `json:"Style_ID"`
	TypeID       *string `json:"Type_ID"`
}

type Company struct {
	ID       *string   `json:"ID"`
	Title    *string   `json:"Title"`
	Verified *bool     `json:"Verified"`
	Owned    *bool     `json:"Owned"`
	Offering *Offering `json:"Offering"`
}

type CompanyInput struct {
	ID       *string   `json:"ID"`
	Title    *string   `json:"Title"`
	Verified *bool     `json:"Verified"`
	Owned    *bool     `json:"Owned"`
	Offering *Offering `json:"Offering"`
}

type Contact struct {
	Email     *string `json:"Email"`
	Phone     *string `json:"Phone"`
	Fname     *string `json:"Fname"`
	Lname     *string `json:"Lname"`
	Twitter   *string `json:"Twitter"`
	Instagram *string `json:"Instagram"`
	TikTok    *string `json:"TikTok"`
	Pinterest *string `json:"Pinterest"`
}

type ContactInput struct {
	Email     *string `json:"Email"`
	Phone     *string `json:"Phone"`
	Fname     *string `json:"Fname"`
	Lname     *string `json:"Lname"`
	Twitter   *string `json:"Twitter"`
	Instagram *string `json:"Instagram"`
	TikTok    *string `json:"TikTok"`
	Pinterest *string `json:"Pinterest"`
	ID        *string `json:"ID"`
}

type Image struct {
	ID  *string `json:"ID"`
	URL *string `json:"URL"`
}

type ImageInput struct {
	ID  *string `json:"ID"`
	URL *string `json:"URL"`
}

type Logo struct {
	ID  *string `json:"ID"`
	URL *string `json:"URL"`
}

type LogoInput struct {
	ID  *string `json:"ID"`
	URL *string `json:"URL"`
}

type Product struct {
	ID              *string  `json:"ID"`
	Title           *string  `json:"Title"`
	Company         *string  `json:"Company"`
	Verified        *bool    `json:"Verified"`
	AverageRating   *float64 `json:"Average_Rating"`
	RatingQuanitity *int     `json:"Rating_Quanitity"`
	PriceBottom     *float64 `json:"Price_Bottom"`
	PriceTop        *float64 `json:"Price_Top"`
}

type ProductInput struct {
	ID              *string  `json:"ID"`
	Title           *string  `json:"Title"`
	Company         *string  `json:"Company"`
	Verified        *bool    `json:"Verified"`
	AverageRating   *float64 `json:"Average_Rating"`
	RatingQuanitity *int     `json:"Rating_Quanitity"`
	PriceBottom     *float64 `json:"Price_Bottom"`
	PriceTop        *float64 `json:"Price_Top"`
}

type ProductNavigation struct {
	ID           *string `json:"ID"`
	ProductID    *string `json:"Product_ID"`
	DepartmentID *string `json:"Department_ID"`
	CategoryID   *string `json:"Category_ID"`
	TypeID       *string `json:"Type_ID"`
	StyleID      *string `json:"Style_ID"`
}

type ProductNavigationInput struct {
	ID           *string `json:"ID"`
	ProductID    *string `json:"Product_ID"`
	DepartmentID *string `json:"Department_ID"`
	CategoryID   *string `json:"Category_ID"`
	TypeID       *string `json:"Type_ID"`
	StyleID      *string `json:"Style_ID"`
}

type Reviews struct {
	ID       *string  `json:"ID"`
	Platform *Website `json:"Platform"`
	Review   *string  `json:"Review"`
	User     *string  `json:"User"`
	Date     *string  `json:"Date"`
	Stars    *float64 `json:"Stars"`
}

type ReviewsInput struct {
	ID       *string  `json:"ID"`
	Platform *Website `json:"Platform"`
	Review   *string  `json:"Review"`
	User     *string  `json:"User"`
	Date     *string  `json:"Date"`
	Stars    *float64 `json:"Stars"`
}

type Style struct {
	ID           *string `json:"ID"`
	Title        *string `json:"Title"`
	DepartmentID *string `json:"Department_ID"`
	CategoryID   *string `json:"Category_ID"`
	TypeID       *string `json:"Type_ID"`
}

type StyleInput struct {
	ID           *string `json:"ID"`
	Title        *string `json:"Title"`
	DepartmentID *string `json:"Department_ID"`
	CategoryID   *string `json:"Category_ID"`
	TypeID       *string `json:"Type_ID"`
}

type Type struct {
	ID           *string `json:"ID"`
	DepartmentID *string `json:"Department_ID"`
	CategoryID   *string `json:"Category_ID"`
	StyleID      *string `json:"Style_ID"`
	Title        *string `json:"Title"`
}

type TypeInput struct {
	ID           *string `json:"ID"`
	DepartmentID *string `json:"Department_ID"`
	CategoryID   *string `json:"Category_ID"`
	StyleID      *string `json:"Style_ID"`
	Title        *string `json:"Title"`
}

type AccessoryType string

const (
	AccessoryTypeCrossBody AccessoryType = "CrossBody"
	AccessoryTypeBackPack  AccessoryType = "BackPack"
	AccessoryTypeShoulder  AccessoryType = "Shoulder"
)

var AllAccessoryType = []AccessoryType{
	AccessoryTypeCrossBody,
	AccessoryTypeBackPack,
	AccessoryTypeShoulder,
}

func (e AccessoryType) IsValid() bool {
	switch e {
	case AccessoryTypeCrossBody, AccessoryTypeBackPack, AccessoryTypeShoulder:
		return true
	}
	return false
}

func (e AccessoryType) String() string {
	return string(e)
}

func (e *AccessoryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccessoryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccessoryType", str)
	}
	return nil
}

func (e AccessoryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AcessoryCategory string

const (
	AcessoryCategoryBracelet  AcessoryCategory = "Bracelet"
	AcessoryCategoryEaring    AcessoryCategory = "Earing"
	AcessoryCategoryRing      AcessoryCategory = "Ring"
	AcessoryCategoryKnecklace AcessoryCategory = "Knecklace"
	AcessoryCategorySet       AcessoryCategory = "Set"
	AcessoryCategoryBelt      AcessoryCategory = "Belt"
	AcessoryCategoryHat       AcessoryCategory = "Hat"
	AcessoryCategoryScarf     AcessoryCategory = "Scarf"
	AcessoryCategoryBag       AcessoryCategory = "Bag"
	AcessoryCategoryWatch     AcessoryCategory = "Watch"
	AcessoryCategoryWallet    AcessoryCategory = "Wallet"
	AcessoryCategoryGloves    AcessoryCategory = "Gloves"
)

var AllAcessoryCategory = []AcessoryCategory{
	AcessoryCategoryBracelet,
	AcessoryCategoryEaring,
	AcessoryCategoryRing,
	AcessoryCategoryKnecklace,
	AcessoryCategorySet,
	AcessoryCategoryBelt,
	AcessoryCategoryHat,
	AcessoryCategoryScarf,
	AcessoryCategoryBag,
	AcessoryCategoryWatch,
	AcessoryCategoryWallet,
	AcessoryCategoryGloves,
}

func (e AcessoryCategory) IsValid() bool {
	switch e {
	case AcessoryCategoryBracelet, AcessoryCategoryEaring, AcessoryCategoryRing, AcessoryCategoryKnecklace, AcessoryCategorySet, AcessoryCategoryBelt, AcessoryCategoryHat, AcessoryCategoryScarf, AcessoryCategoryBag, AcessoryCategoryWatch, AcessoryCategoryWallet, AcessoryCategoryGloves:
		return true
	}
	return false
}

func (e AcessoryCategory) String() string {
	return string(e)
}

func (e *AcessoryCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AcessoryCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AcessoryCategory", str)
	}
	return nil
}

func (e AcessoryCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ClothingCategory string

const (
	ClothingCategoryMenIdentifying   ClothingCategory = "MenIdentifying"
	ClothingCategoryWomenIdentifying ClothingCategory = "WomenIdentifying"
	ClothingCategoryUnisex           ClothingCategory = "Unisex"
	ClothingCategoryToddler          ClothingCategory = "Toddler"
	ClothingCategoryBaby             ClothingCategory = "Baby"
	ClothingCategoryPetite           ClothingCategory = "Petite"
	ClothingCategoryPlus             ClothingCategory = "Plus"
	ClothingCategoryMaternity        ClothingCategory = "Maternity"
	ClothingCategoryChildren         ClothingCategory = "Children"
)

var AllClothingCategory = []ClothingCategory{
	ClothingCategoryMenIdentifying,
	ClothingCategoryWomenIdentifying,
	ClothingCategoryUnisex,
	ClothingCategoryToddler,
	ClothingCategoryBaby,
	ClothingCategoryPetite,
	ClothingCategoryPlus,
	ClothingCategoryMaternity,
	ClothingCategoryChildren,
}

func (e ClothingCategory) IsValid() bool {
	switch e {
	case ClothingCategoryMenIdentifying, ClothingCategoryWomenIdentifying, ClothingCategoryUnisex, ClothingCategoryToddler, ClothingCategoryBaby, ClothingCategoryPetite, ClothingCategoryPlus, ClothingCategoryMaternity, ClothingCategoryChildren:
		return true
	}
	return false
}

func (e ClothingCategory) String() string {
	return string(e)
}

func (e *ClothingCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ClothingCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ClothingCategory", str)
	}
	return nil
}

func (e ClothingCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ClothingType string

const (
	ClothingTypeTops       ClothingType = "Tops"
	ClothingTypeSweaters   ClothingType = "Sweaters"
	ClothingTypeJeans      ClothingType = "Jeans"
	ClothingTypeCoats      ClothingType = "Coats"
	ClothingTypePants      ClothingType = "Pants"
	ClothingTypeShorts     ClothingType = "Shorts"
	ClothingTypeSkirts     ClothingType = "Skirts"
	ClothingTypeSwim       ClothingType = "Swim"
	ClothingTypeActivewear ClothingType = "Activewear"
	ClothingTypeSkorts     ClothingType = "Skorts"
	ClothingTypeLingerie   ClothingType = "Lingerie"
	ClothingTypeDresses    ClothingType = "Dresses"
)

var AllClothingType = []ClothingType{
	ClothingTypeTops,
	ClothingTypeSweaters,
	ClothingTypeJeans,
	ClothingTypeCoats,
	ClothingTypePants,
	ClothingTypeShorts,
	ClothingTypeSkirts,
	ClothingTypeSwim,
	ClothingTypeActivewear,
	ClothingTypeSkorts,
	ClothingTypeLingerie,
	ClothingTypeDresses,
}

func (e ClothingType) IsValid() bool {
	switch e {
	case ClothingTypeTops, ClothingTypeSweaters, ClothingTypeJeans, ClothingTypeCoats, ClothingTypePants, ClothingTypeShorts, ClothingTypeSkirts, ClothingTypeSwim, ClothingTypeActivewear, ClothingTypeSkorts, ClothingTypeLingerie, ClothingTypeDresses:
		return true
	}
	return false
}

func (e ClothingType) String() string {
	return string(e)
}

func (e *ClothingType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ClothingType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ClothingType", str)
	}
	return nil
}

func (e ClothingType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Color string

const (
	ColorBlue       Color = "Blue"
	ColorRed        Color = "Red"
	ColorOrange     Color = "Orange"
	ColorPink       Color = "Pink"
	ColorPurple     Color = "Purple"
	ColorYellow     Color = "Yellow"
	ColorTan        Color = "Tan"
	ColorBrown      Color = "Brown"
	ColorGreen      Color = "Green"
	ColorWhite      Color = "White"
	ColorGrey       Color = "Grey"
	ColorBlack      Color = "Black"
	ColorSilver     Color = "Silver"
	ColorGold       Color = "Gold"
	ColorYellowGold Color = "YellowGold"
	ColorWhiteGold  Color = "WhiteGold"
)

var AllColor = []Color{
	ColorBlue,
	ColorRed,
	ColorOrange,
	ColorPink,
	ColorPurple,
	ColorYellow,
	ColorTan,
	ColorBrown,
	ColorGreen,
	ColorWhite,
	ColorGrey,
	ColorBlack,
	ColorSilver,
	ColorGold,
	ColorYellowGold,
	ColorWhiteGold,
}

func (e Color) IsValid() bool {
	switch e {
	case ColorBlue, ColorRed, ColorOrange, ColorPink, ColorPurple, ColorYellow, ColorTan, ColorBrown, ColorGreen, ColorWhite, ColorGrey, ColorBlack, ColorSilver, ColorGold, ColorYellowGold, ColorWhiteGold:
		return true
	}
	return false
}

func (e Color) String() string {
	return string(e)
}

func (e *Color) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Color(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Color", str)
	}
	return nil
}

func (e Color) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Country string

const (
	CountryUnitedStates Country = "United_States"
)

var AllCountry = []Country{
	CountryUnitedStates,
}

func (e Country) IsValid() bool {
	switch e {
	case CountryUnitedStates:
		return true
	}
	return false
}

func (e Country) String() string {
	return string(e)
}

func (e *Country) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Country(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Country", str)
	}
	return nil
}

func (e Country) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Department string

const (
	// Clothing and shoes
	DepartmentClothing    Department = "Clothing"
	DepartmentJewelry     Department = "Jewelry"
	DepartmentAccessories Department = "Accessories"
	DepartmentHome        Department = "Home"
	DepartmentGarden      Department = "Garden"
	DepartmentShoes       Department = "Shoes"
	DepartmentBeauty      Department = "Beauty"
	DepartmentBaby        Department = "Baby"
	DepartmentGrocery     Department = "Grocery"
	DepartmentPets        Department = "Pets"
)

var AllDepartment = []Department{
	DepartmentClothing,
	DepartmentJewelry,
	DepartmentAccessories,
	DepartmentHome,
	DepartmentGarden,
	DepartmentShoes,
	DepartmentBeauty,
	DepartmentBaby,
	DepartmentGrocery,
	DepartmentPets,
}

func (e Department) IsValid() bool {
	switch e {
	case DepartmentClothing, DepartmentJewelry, DepartmentAccessories, DepartmentHome, DepartmentGarden, DepartmentShoes, DepartmentBeauty, DepartmentBaby, DepartmentGrocery, DepartmentPets:
		return true
	}
	return false
}

func (e Department) String() string {
	return string(e)
}

func (e *Department) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Department(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Department", str)
	}
	return nil
}

func (e Department) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Material string

const (
	MaterialCotton  Material = "Cotton"
	MaterialSpandex Material = "Spandex"
	MaterialLycra   Material = "Lycra"
	MaterialLinen   Material = "Linen"
)

var AllMaterial = []Material{
	MaterialCotton,
	MaterialSpandex,
	MaterialLycra,
	MaterialLinen,
}

func (e Material) IsValid() bool {
	switch e {
	case MaterialCotton, MaterialSpandex, MaterialLycra, MaterialLinen:
		return true
	}
	return false
}

func (e Material) String() string {
	return string(e)
}

func (e *Material) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Material(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Material", str)
	}
	return nil
}

func (e Material) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Offering string

const (
	OfferingService Offering = "Service"
	OfferingProduct Offering = "Product"
)

var AllOffering = []Offering{
	OfferingService,
	OfferingProduct,
}

func (e Offering) IsValid() bool {
	switch e {
	case OfferingService, OfferingProduct:
		return true
	}
	return false
}

func (e Offering) String() string {
	return string(e)
}

func (e *Offering) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Offering(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Offering", str)
	}
	return nil
}

func (e Offering) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShoeSize string

const (
	ShoeSizeFour     ShoeSize = "Four"
	ShoeSizeFive     ShoeSize = "Five"
	ShoeSizeSix      ShoeSize = "Six"
	ShoeSizeSeven    ShoeSize = "Seven"
	ShoeSizeEight    ShoeSize = "Eight"
	ShoeSizeNine     ShoeSize = "Nine"
	ShoeSizeTen      ShoeSize = "Ten"
	ShoeSizeEleven   ShoeSize = "Eleven"
	ShoeSizeTwelve   ShoeSize = "Twelve"
	ShoeSizeThirteen ShoeSize = "Thirteen"
)

var AllShoeSize = []ShoeSize{
	ShoeSizeFour,
	ShoeSizeFive,
	ShoeSizeSix,
	ShoeSizeSeven,
	ShoeSizeEight,
	ShoeSizeNine,
	ShoeSizeTen,
	ShoeSizeEleven,
	ShoeSizeTwelve,
	ShoeSizeThirteen,
}

func (e ShoeSize) IsValid() bool {
	switch e {
	case ShoeSizeFour, ShoeSizeFive, ShoeSizeSix, ShoeSizeSeven, ShoeSizeEight, ShoeSizeNine, ShoeSizeTen, ShoeSizeEleven, ShoeSizeTwelve, ShoeSizeThirteen:
		return true
	}
	return false
}

func (e ShoeSize) String() string {
	return string(e)
}

func (e *ShoeSize) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShoeSize(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShoeSize", str)
	}
	return nil
}

func (e ShoeSize) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShoesType string

const (
	ShoesTypeBoots          ShoesType = "Boots"
	ShoesTypeSandal         ShoesType = "Sandal"
	ShoesTypeSlipper        ShoesType = "Slipper"
	ShoesTypePump           ShoesType = "Pump"
	ShoesTypeFashionSneaker ShoesType = "FashionSneaker"
	ShoesTypeAthletic       ShoesType = "Athletic"
	ShoesTypeFlats          ShoesType = "Flats"
)

var AllShoesType = []ShoesType{
	ShoesTypeBoots,
	ShoesTypeSandal,
	ShoesTypeSlipper,
	ShoesTypePump,
	ShoesTypeFashionSneaker,
	ShoesTypeAthletic,
	ShoesTypeFlats,
}

func (e ShoesType) IsValid() bool {
	switch e {
	case ShoesTypeBoots, ShoesTypeSandal, ShoesTypeSlipper, ShoesTypePump, ShoesTypeFashionSneaker, ShoesTypeAthletic, ShoesTypeFlats:
		return true
	}
	return false
}

func (e ShoesType) String() string {
	return string(e)
}

func (e *ShoesType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShoesType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShoesType", str)
	}
	return nil
}

func (e ShoesType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SizeRange string

const (
	SizeRangePetite    SizeRange = "Petite"
	SizeRangeRegular   SizeRange = "Regular"
	SizeRangeJunior    SizeRange = "Junior"
	SizeRangeMaternity SizeRange = "Maternity"
)

var AllSizeRange = []SizeRange{
	SizeRangePetite,
	SizeRangeRegular,
	SizeRangeJunior,
	SizeRangeMaternity,
}

func (e SizeRange) IsValid() bool {
	switch e {
	case SizeRangePetite, SizeRangeRegular, SizeRangeJunior, SizeRangeMaternity:
		return true
	}
	return false
}

func (e SizeRange) String() string {
	return string(e)
}

func (e *SizeRange) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SizeRange(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SizeRange", str)
	}
	return nil
}

func (e SizeRange) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type USStates string

const (
	USStatesWyoming       USStates = "Wyoming"
	USStatesWisconsin     USStates = "Wisconsin"
	USStatesWestVirginia  USStates = "West_Virginia"
	USStatesWashington    USStates = "Washington"
	USStatesVirginia1     USStates = "Virginia1"
	USStatesVermont       USStates = "Vermont"
	USStatesUtah          USStates = "Utah"
	USStatesTexas         USStates = "Texas"
	USStatesTennessee     USStates = "Tennessee"
	USStatesSouthDakota   USStates = "South_Dakota"
	USStatesSouthCarolina USStates = "South_Carolina"
	USStatesRhodeIsland   USStates = "Rhode_Island"
	USStatesPennsylvania  USStates = "Pennsylvania"
	USStatesOregon        USStates = "Oregon"
	USStatesOklahoma      USStates = "Oklahoma"
	USStatesOhio          USStates = "Ohio"
	USStatesNorthDakota   USStates = "North_Dakota"
	USStatesNorthCarolina USStates = "North_Carolina"
	USStatesNewYork       USStates = "New_York"
	USStatesNewMexico     USStates = "New_Mexico"
	USStatesNewJersey     USStates = "New_Jersey"
	USStatesNewHampshire  USStates = "New_Hampshire"
	USStatesNevada        USStates = "Nevada"
	USStatesNebraska      USStates = "Nebraska"
	USStatesMontana       USStates = "Montana"
	USStatesMissouri      USStates = "Missouri"
	USStatesMississippi   USStates = "Mississippi"
	USStatesMinnesota     USStates = "Minnesota"
	USStatesMichigan      USStates = "Michigan"
	USStatesMassachusetts USStates = "Massachusetts"
	USStatesMaryland      USStates = "Maryland"
	USStatesMaine         USStates = "Maine"
	USStatesLouisiana     USStates = "Louisiana"
	USStatesKentucky      USStates = "Kentucky"
	USStatesKansas        USStates = "Kansas"
	USStatesIowa          USStates = "Iowa"
	USStatesIndiana       USStates = "Indiana"
	USStatesIllinois      USStates = "Illinois"
	USStatesIdaho         USStates = "Idaho"
	USStatesHawaii        USStates = "Hawaii"
	USStatesGeorgia       USStates = "Georgia"
	USStatesFlorida       USStates = "Florida"
	USStatesDelaware      USStates = "Delaware"
	USStatesConnecticut   USStates = "Connecticut"
	USStatesColorado      USStates = "Colorado"
	USStatesCalifornia    USStates = "California"
	USStatesArkansas      USStates = "Arkansas"
	USStatesArizona       USStates = "Arizona"
	USStatesAlaska        USStates = "Alaska"
	USStatesAlabama       USStates = "Alabama"
)

var AllUSStates = []USStates{
	USStatesWyoming,
	USStatesWisconsin,
	USStatesWestVirginia,
	USStatesWashington,
	USStatesVirginia1,
	USStatesVermont,
	USStatesUtah,
	USStatesTexas,
	USStatesTennessee,
	USStatesSouthDakota,
	USStatesSouthCarolina,
	USStatesRhodeIsland,
	USStatesPennsylvania,
	USStatesOregon,
	USStatesOklahoma,
	USStatesOhio,
	USStatesNorthDakota,
	USStatesNorthCarolina,
	USStatesNewYork,
	USStatesNewMexico,
	USStatesNewJersey,
	USStatesNewHampshire,
	USStatesNevada,
	USStatesNebraska,
	USStatesMontana,
	USStatesMissouri,
	USStatesMississippi,
	USStatesMinnesota,
	USStatesMichigan,
	USStatesMassachusetts,
	USStatesMaryland,
	USStatesMaine,
	USStatesLouisiana,
	USStatesKentucky,
	USStatesKansas,
	USStatesIowa,
	USStatesIndiana,
	USStatesIllinois,
	USStatesIdaho,
	USStatesHawaii,
	USStatesGeorgia,
	USStatesFlorida,
	USStatesDelaware,
	USStatesConnecticut,
	USStatesColorado,
	USStatesCalifornia,
	USStatesArkansas,
	USStatesArizona,
	USStatesAlaska,
	USStatesAlabama,
}

func (e USStates) IsValid() bool {
	switch e {
	case USStatesWyoming, USStatesWisconsin, USStatesWestVirginia, USStatesWashington, USStatesVirginia1, USStatesVermont, USStatesUtah, USStatesTexas, USStatesTennessee, USStatesSouthDakota, USStatesSouthCarolina, USStatesRhodeIsland, USStatesPennsylvania, USStatesOregon, USStatesOklahoma, USStatesOhio, USStatesNorthDakota, USStatesNorthCarolina, USStatesNewYork, USStatesNewMexico, USStatesNewJersey, USStatesNewHampshire, USStatesNevada, USStatesNebraska, USStatesMontana, USStatesMissouri, USStatesMississippi, USStatesMinnesota, USStatesMichigan, USStatesMassachusetts, USStatesMaryland, USStatesMaine, USStatesLouisiana, USStatesKentucky, USStatesKansas, USStatesIowa, USStatesIndiana, USStatesIllinois, USStatesIdaho, USStatesHawaii, USStatesGeorgia, USStatesFlorida, USStatesDelaware, USStatesConnecticut, USStatesColorado, USStatesCalifornia, USStatesArkansas, USStatesArizona, USStatesAlaska, USStatesAlabama:
		return true
	}
	return false
}

func (e USStates) String() string {
	return string(e)
}

func (e *USStates) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = USStates(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid USStates", str)
	}
	return nil
}

func (e USStates) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Website string

const (
	WebsiteAmazon      Website = "Amazon"
	WebsiteEtsy        Website = "Etsy"
	WebsiteShopify     Website = "Shopify"
	WebsiteWooCommerce Website = "WooCommerce"
)

var AllWebsite = []Website{
	WebsiteAmazon,
	WebsiteEtsy,
	WebsiteShopify,
	WebsiteWooCommerce,
}

func (e Website) IsValid() bool {
	switch e {
	case WebsiteAmazon, WebsiteEtsy, WebsiteShopify, WebsiteWooCommerce:
		return true
	}
	return false
}

func (e Website) String() string {
	return string(e)
}

func (e *Website) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Website(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Website", str)
	}
	return nil
}

func (e Website) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}