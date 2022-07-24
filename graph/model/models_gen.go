// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type ArticleConnection struct {
	Edges    []*ArticleEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type ArticleEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Article `json:"node"`
}

type ArticleFilter struct {
	Title *string `json:"title"`
}

type ArticleInput struct {
	Description string   `json:"description"`
	FileName    *string  `json:"fileName"`
	Published   bool     `json:"published"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
}

type ArticleInputWithID struct {
	ID    string        `json:"id"`
	Input *ArticleInput `json:"input"`
}

type ArticlePayload struct {
	Record   *Article `json:"record"`
	RecordID string   `json:"recordId"`
}

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

type Cart struct {
	ID        string   `json:"id"`
	KitIds    []string `json:"kitIds"`
	Published bool     `json:"published"`
	Student   *Student `json:"student"`
	StudentID string   `json:"studentId"`
	Sum       int      `json:"sum"`
}

type CartConnection struct {
	Edges    []*CartEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type CartDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CartEdge struct {
	Cursor string `json:"cursor"`
	Node   *Cart  `json:"node"`
}

type CartFilter struct {
	Name *string `json:"name"`
}

type CartInput struct {
	KitIds    []string `json:"kitIds"`
	Published bool     `json:"published"`
	StudentID string   `json:"studentId"`
	Sum       int      `json:"sum"`
}

type CartInputWithID struct {
	ID    string     `json:"id"`
	Input *CartInput `json:"input"`
}

type CartPayload struct {
	Record   *Cart  `json:"record"`
	RecordID string `json:"recordId"`
}

type ClubBalanceConnection struct {
	Edges    []*ClubBalanceEdge `json:"edges"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type ClubBalanceEdge struct {
	Cursor string       `json:"cursor"`
	Node   *ClubBalance `json:"node"`
}

type ClubBalanceFilter struct {
	Name *string `json:"name"`
}

type ClubBalanceInput struct {
	Date       time.Time `json:"date"`
	OtherCosts int       `json:"otherCosts"`
	Published  bool      `json:"published"`
	Rent       int       `json:"rent"`
	Salary     int       `json:"salary"`
	Tickets    int       `json:"tickets"`
}

type ClubBalanceInputWithID struct {
	ID    string            `json:"id"`
	Input *ClubBalanceInput `json:"input"`
}

type ClubBalancePayload struct {
	Record   *ClubBalance `json:"record"`
	RecordID string       `json:"recordId"`
}

type CoachPaymentByMonthConnection struct {
	Edges    []*CoachPaymentByMonthEdge `json:"edges"`
	PageInfo *PageInfo                  `json:"pageInfo"`
}

type CoachPaymentByMonthEdge struct {
	Cursor string               `json:"cursor"`
	Node   *CoachPaymentByMonth `json:"node"`
}

type CoachPaymentByMonthFilter struct {
	Name *string `json:"name"`
}

type CoachPaymentByMonthInput struct {
	CoachID   string    `json:"coachId"`
	Date      time.Time `json:"date"`
	Published bool      `json:"published"`
	Sum       int       `json:"sum"`
}

type CoachPaymentByMonthInputWithID struct {
	ID    string                    `json:"id"`
	Input *CoachPaymentByMonthInput `json:"input"`
}

type CoachPaymentByMonthPayload struct {
	Record   *CoachPaymentByMonth `json:"record"`
	RecordID string               `json:"recordId"`
}

type CoachPaymentByTeamConnection struct {
	Edges    []*CoachPaymentByTeamEdge `json:"edges"`
	PageInfo *PageInfo                 `json:"pageInfo"`
}

type CoachPaymentByTeamEdge struct {
	Cursor string              `json:"cursor"`
	Node   *CoachPaymentByTeam `json:"node"`
}

type CoachPaymentByTeamFilter struct {
	Name *string `json:"name"`
}

type CoachPaymentByTeamInput struct {
	CoachID     string           `json:"coachId"`
	DateFinish  time.Time        `json:"dateFinish"`
	DateStart   time.Time        `json:"dateStart"`
	PaymentRule CoachPaymentRule `json:"paymentRule"`
	Published   bool             `json:"published"`
	Sum         int              `json:"sum"`
	TeamID      string           `json:"teamID"`
}

type CoachPaymentByTeamInputWithID struct {
	ID    string                   `json:"id"`
	Input *CoachPaymentByTeamInput `json:"input"`
}

type CoachPaymentByTeamPayload struct {
	Record   *CoachPaymentByTeam `json:"record"`
	RecordID string              `json:"recordId"`
}

type CoachPaymentByTrainingConnection struct {
	Edges    []*CoachPaymentByTrainingEdge `json:"edges"`
	PageInfo *PageInfo                     `json:"pageInfo"`
}

type CoachPaymentByTrainingEdge struct {
	Cursor string                  `json:"cursor"`
	Node   *CoachPaymentByTraining `json:"node"`
}

type CoachPaymentByTrainingFilter struct {
	Name *string `json:"name"`
}

type CoachPaymentByTrainingInput struct {
	CoachID    string `json:"coachId"`
	Published  bool   `json:"published"`
	Sum        int    `json:"sum"`
	TrainingID string `json:"trainingID"`
}

type CoachPaymentByTrainingInputWithID struct {
	ID    string                       `json:"id"`
	Input *CoachPaymentByTrainingInput `json:"input"`
}

type CoachPaymentByTrainingPayload struct {
	Record   *CoachPaymentByTraining `json:"record"`
	RecordID string                  `json:"recordId"`
}

type CreatorConnection struct {
	Edges    []*CreatorEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type CreatorDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatorEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Creator `json:"node"`
}

type CreatorFilter struct {
	Name *string `json:"name"`
}

type CreatorInput struct {
	Name        string  `json:"name"`
	PassportNum *string `json:"passportNum"`
	Phone       string  `json:"phone"`
	Published   bool    `json:"published"`
	UserID      *string `json:"userId"`
}

type CreatorInputWithID struct {
	ID    string        `json:"id"`
	Input *CreatorInput `json:"input"`
}

type CreatorPayload struct {
	Record   *Creator `json:"record"`
	RecordID string   `json:"recordId"`
}

type Kit struct {
	FileName  *string `json:"fileName"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Number    *int    `json:"number"`
	Price     int     `json:"price"`
	Published bool    `json:"published"`
	Quantity  *int    `json:"quantity"`
	Size      string  `json:"size"`
	Title     *string `json:"title"`
}

type KitConnection struct {
	Edges    []*KitEdge `json:"edges"`
	PageInfo *PageInfo  `json:"pageInfo"`
}

type KitDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type KitEdge struct {
	Cursor string `json:"cursor"`
	Node   *Kit   `json:"node"`
}

type KitFilter struct {
	Name *string `json:"name"`
}

type KitInput struct {
	FileName  *string `json:"fileName"`
	Name      string  `json:"name"`
	Number    *int    `json:"number"`
	Price     int     `json:"price"`
	Published bool    `json:"published"`
	Quantity  *int    `json:"quantity"`
	Size      string  `json:"size"`
	Title     *string `json:"title"`
}

type KitInputWithID struct {
	ID    string    `json:"id"`
	Input *KitInput `json:"input"`
}

type KitPayload struct {
	Record   *Kit   `json:"record"`
	RecordID string `json:"recordId"`
}

type Lead struct {
	Description *string     `json:"description"`
	ID          string      `json:"id"`
	Name        *string     `json:"name"`
	NextVisit   *Training   `json:"nextVisit"`
	NextVisitID *string     `json:"nextVisitId"`
	Phone       string      `json:"phone"`
	Published   bool        `json:"published"`
	Source      *LeadSource `json:"source"`
	Status      *LeadStatus `json:"status"`
	StudentIds  []string    `json:"studentIds"`
	Students    []*Student  `json:"students"`
	Team        *Team       `json:"team"`
	TeamID      *string     `json:"teamId"`
	YearBorn    *int        `json:"yearBorn"`
}

type LeadConnection struct {
	Edges    []*LeadEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type LeadDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LeadEdge struct {
	Cursor string `json:"cursor"`
	Node   *Lead  `json:"node"`
}

type LeadFilter struct {
	Name *string `json:"name"`
}

type LeadInput struct {
	Description *string     `json:"description"`
	Name        *string     `json:"name"`
	NextVisitID *string     `json:"nextVisitId"`
	Phone       string      `json:"phone"`
	Published   bool        `json:"published"`
	Source      *LeadSource `json:"source"`
	Status      *LeadStatus `json:"status"`
	StudentIds  []string    `json:"studentIds"`
	TeamID      *string     `json:"teamId"`
	YearBorn    *int        `json:"yearBorn"`
}

type LeadInputWithID struct {
	ID    string     `json:"id"`
	Input *LeadInput `json:"input"`
}

type LeadPayload struct {
	Record   *Lead  `json:"record"`
	RecordID string `json:"recordId"`
}

type LoginInput struct {
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type MoneyCostConnection struct {
	Edges    []*MoneyCostEdge `json:"edges"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type MoneyCostEdge struct {
	Cursor string     `json:"cursor"`
	Node   *MoneyCost `json:"node"`
}

type MoneyCostFilter struct {
	Name *string `json:"name"`
}

type MoneyCostInput struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	MoneyForm   MoneyForm `json:"moneyForm"`
	Published   bool      `json:"published"`
	StaffID     string    `json:"staffId"`
	Sum         int       `json:"sum"`
}

type MoneyCostInputWithID struct {
	ID    string          `json:"id"`
	Input *MoneyCostInput `json:"input"`
}

type MoneyCostPayload struct {
	Record   *MoneyCost `json:"record"`
	RecordID string     `json:"recordId"`
}

type MoneyMoveConnection struct {
	Edges    []*MoneyMoveEdge `json:"edges"`
	PageInfo *PageInfo        `json:"pageInfo"`
}

type MoneyMoveEdge struct {
	Cursor string     `json:"cursor"`
	Node   *MoneyMove `json:"node"`
}

type MoneyMoveFilter struct {
	Name *string `json:"name"`
}

type MoneyMoveInput struct {
	DateFinish  time.Time `json:"dateFinish"`
	DatePayment time.Time `json:"datePayment"`
	DateStart   time.Time `json:"dateStart"`
	Description *string   `json:"description"`
	MoneyForm   MoneyForm `json:"moneyForm"`
	OwnerID     string    `json:"ownerId"`
	Published   bool      `json:"published"`
	StudentID   string    `json:"studentId"`
	Sum         int       `json:"sum"`
	UserID      string    `json:"userId"`
}

type MoneyMoveInputWithID struct {
	ID    string          `json:"id"`
	Input *MoneyMoveInput `json:"input"`
}

type MoneyMovePayload struct {
	Record   *MoneyMove `json:"record"`
	RecordID string     `json:"recordId"`
}

type Order struct {
	Cart        *Cart       `json:"cart"`
	CartID      string      `json:"cartId"`
	Creator     *Creator    `json:"creator"`
	CreatorID   string      `json:"creatorId"`
	FileName    *string     `json:"fileName"`
	ID          string      `json:"id"`
	OrderStatus OrderStatus `json:"orderStatus"`
	Published   bool        `json:"published"`
}

type OrderConnection struct {
	Edges    []*OrderEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

type OrderEdge struct {
	Cursor string `json:"cursor"`
	Node   *Order `json:"node"`
}

type OrderFilter struct {
	Name *string `json:"name"`
}

type OrderInput struct {
	CartID      string      `json:"cartId"`
	CreatorID   string      `json:"creatorId"`
	FileName    *string     `json:"fileName"`
	OrderStatus OrderStatus `json:"orderStatus"`
	Published   bool        `json:"published"`
}

type OrderInputWithID struct {
	ID    string      `json:"id"`
	Input *OrderInput `json:"input"`
}

type OrderPayload struct {
	Record   *Order `json:"record"`
	RecordID string `json:"recordId"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage *bool  `json:"hasNextPage"`
	StartCursor string `json:"startCursor"`
}

type PlaceConnection struct {
	Edges    []*PlaceEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

type PlaceDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PlaceEdge struct {
	Cursor string `json:"cursor"`
	Node   *Place `json:"node"`
}

type PlaceFilter struct {
	Name *string `json:"name"`
}

type PlaceInput struct {
	Address     string `json:"address"`
	Description string `json:"description"`
	Name        string `json:"name"`
	OrderNumber int    `json:"orderNumber"`
	Published   bool   `json:"published"`
}

type PlaceInputWithID struct {
	ID    string      `json:"id"`
	Input *PlaceInput `json:"input"`
}

type PlacePayload struct {
	Record   *Place `json:"record"`
	RecordID string `json:"recordId"`
}

type RegisterInput struct {
	ConfirmPassword string `json:"confirmPassword"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	Phone           string `json:"phone"`
}

type RentPaymentByMonthConnection struct {
	Edges    []*RentPaymentByMonthEdge `json:"edges"`
	PageInfo *PageInfo                 `json:"pageInfo"`
}

type RentPaymentByMonthEdge struct {
	Cursor string              `json:"cursor"`
	Node   *RentPaymentByMonth `json:"node"`
}

type RentPaymentByMonthFilter struct {
	Name *string `json:"name"`
}

type RentPaymentByMonthInput struct {
	Description *string   `json:"description"`
	Month       time.Time `json:"month"`
	PaymentDate time.Time `json:"paymentDate"`
	Published   bool      `json:"published"`
	StadiumID   string    `json:"stadiumId"`
	Sum         int       `json:"sum"`
}

type RentPaymentByMonthInputWithID struct {
	ID    string                   `json:"id"`
	Input *RentPaymentByMonthInput `json:"input"`
}

type RentPaymentByMonthPayload struct {
	Record   *RentPaymentByMonth `json:"record"`
	RecordID string              `json:"recordId"`
}

type RentPaymentByTrainingConnection struct {
	Edges    []*RentPaymentByTrainingEdge `json:"edges"`
	PageInfo *PageInfo                    `json:"pageInfo"`
}

type RentPaymentByTrainingEdge struct {
	Cursor string                 `json:"cursor"`
	Node   *RentPaymentByTraining `json:"node"`
}

type RentPaymentByTrainingFilter struct {
	Name *string `json:"name"`
}

type RentPaymentByTrainingInput struct {
	Description *string  `json:"description"`
	Published   bool     `json:"published"`
	StadiumID   string   `json:"stadiumID"`
	Sum         int      `json:"sum"`
	TrainingIds []string `json:"trainingIds"`
}

type RentPaymentByTrainingInputWithID struct {
	ID    string                      `json:"id"`
	Input *RentPaymentByTrainingInput `json:"input"`
}

type RentPaymentByTrainingPayload struct {
	Record   *RentPaymentByTraining `json:"record"`
	RecordID string                 `json:"recordId"`
}

type StadiumConnection struct {
	Edges    []*StadiumEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type StadiumDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StadiumEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Stadium `json:"node"`
}

type StadiumFilter struct {
	Name *string `json:"name"`
}

type StadiumInput struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name"`
	PlaceID   string  `json:"placeId"`
	Published bool    `json:"published"`
}

type StadiumInputWithID struct {
	ID    string        `json:"id"`
	Input *StadiumInput `json:"input"`
}

type StadiumPayload struct {
	Record   *Stadium `json:"record"`
	RecordID string   `json:"recordId"`
}

type StaffConnection struct {
	Edges    []*StaffEdge `json:"edges"`
	PageInfo *PageInfo    `json:"pageInfo"`
}

type StaffDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StaffEdge struct {
	Cursor string `json:"cursor"`
	Node   *Staff `json:"node"`
}

type StaffFilter struct {
	Name *string `json:"name"`
}

type StaffInput struct {
	Birthday    *time.Time `json:"birthday"`
	Department  Department `json:"department"`
	Description *string    `json:"description"`
	FileName    *string    `json:"fileName"`
	Name        string     `json:"name"`
	OrderNumber int        `json:"orderNumber"`
	PhoneNumber *string    `json:"phoneNumber"`
	Published   bool       `json:"published"`
	UserID      *string    `json:"userId"`
	Work        string     `json:"work"`
}

type StaffInputWithID struct {
	ID    string      `json:"id"`
	Input *StaffInput `json:"input"`
}

type StaffPayload struct {
	Record   *Staff `json:"record"`
	RecordID string `json:"recordId"`
}

type StudentConnection struct {
	Edges    []*StudentEdge `json:"edges"`
	PageInfo *PageInfo      `json:"pageInfo"`
}

type StudentDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StudentEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Student `json:"node"`
}

type StudentFilter struct {
	Name *string `json:"name"`
}

type StudentInput struct {
	Birthday    *time.Time `json:"birthday"`
	CreatorIds  []string   `json:"creatorIds"`
	Name        string     `json:"name"`
	PassportNum *string    `json:"passportNum"`
	PaymentSum  *int       `json:"paymentSum"`
	Published   bool       `json:"published"`
	TeamIds     []string   `json:"teamIds"`
}

type StudentInputWithID struct {
	ID    string        `json:"id"`
	Input *StudentInput `json:"input"`
}

type StudentPayload struct {
	Record   *Student `json:"record"`
	RecordID string   `json:"recordId"`
}

type StudentVisitConnection struct {
	Edges    []*StudentVisitEdge `json:"edges"`
	PageInfo *PageInfo           `json:"pageInfo"`
}

type StudentVisitEdge struct {
	Cursor string        `json:"cursor"`
	Node   *StudentVisit `json:"node"`
}

type StudentVisitFilter struct {
	Name *string `json:"name"`
}

type StudentVisitInput struct {
	Payed       bool        `json:"payed"`
	Published   bool        `json:"published"`
	StudentID   string      `json:"studentId"`
	TrainingID  string      `json:"trainingId"`
	VisitStatus VisitStatus `json:"visitStatus"`
}

type StudentVisitInputWithID struct {
	ID    string             `json:"id"`
	Input *StudentVisitInput `json:"input"`
}

type StudentVisitPayload struct {
	Record   *StudentVisit `json:"record"`
	RecordID string        `json:"recordId"`
}

type Task struct {
	Author      *User       `json:"author"`
	AuthorID    *string     `json:"authorId"`
	Description *string     `json:"description"`
	EndTime     *time.Time  `json:"endTime"`
	ID          string      `json:"id"`
	LeadIds     []string    `json:"leadIds"`
	Leads       []*Lead     `json:"leads"`
	Priority    *Priority   `json:"priority"`
	Published   bool        `json:"published"`
	Result      *string     `json:"result"`
	StartTime   *time.Time  `json:"startTime"`
	StudentIds  []string    `json:"studentIds"`
	Students    []*Student  `json:"students"`
	TaskStatus  *TaskStatus `json:"taskStatus"`
	Title       string      `json:"title"`
	WorkerIds   []string    `json:"workerIds"`
	Workers     []*Staff    `json:"workers"`
}

type TaskConnection struct {
	Edges    []*TaskEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TaskEdge struct {
	Cursor string `json:"cursor"`
	Node   *Task  `json:"node"`
}

type TaskFilter struct {
	Name *string `json:"name"`
}

type TaskInput struct {
	AuthorID    *string     `json:"authorId"`
	Description *string     `json:"description"`
	EndTime     *time.Time  `json:"endTime"`
	LeadIds     []string    `json:"leadIds"`
	Priority    *Priority   `json:"priority"`
	Published   bool        `json:"published"`
	Result      *string     `json:"result"`
	StartTime   *time.Time  `json:"startTime"`
	StudentIds  []string    `json:"studentIds"`
	TaskStatus  *TaskStatus `json:"taskStatus"`
	Title       string      `json:"title"`
	WorkerIds   []string    `json:"workerIds"`
}

type TaskInputWithID struct {
	ID    string     `json:"id"`
	Input *TaskInput `json:"input"`
}

type TaskPayload struct {
	Record   *Task  `json:"record"`
	RecordID string `json:"recordId"`
}

type TeamBalanceConnection struct {
	Edges    []*TeamBalanceEdge `json:"edges"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type TeamBalanceEdge struct {
	Cursor string       `json:"cursor"`
	Node   *TeamBalance `json:"node"`
}

type TeamBalanceFilter struct {
	Name *string `json:"name"`
}

type TeamBalanceInput struct {
	Date      time.Time `json:"date"`
	Published bool      `json:"published"`
	Rent      int       `json:"rent"`
	Salary    int       `json:"salary"`
	TeamID    string    `json:"teamId"`
	Tickets   int       `json:"tickets"`
}

type TeamBalanceInputWithID struct {
	ID    string            `json:"id"`
	Input *TeamBalanceInput `json:"input"`
}

type TeamBalancePayload struct {
	Record   *TeamBalance `json:"record"`
	RecordID string       `json:"recordId"`
}

type TeamConnection struct {
	Edges    []*TeamEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type TeamDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TeamEdge struct {
	Cursor string `json:"cursor"`
	Node   *Team  `json:"node"`
}

type TeamFilter struct {
	Name *string `json:"name"`
}

type TeamInput struct {
	Ages        []Age    `json:"ages"`
	CoachIds    []string `json:"coachIds"`
	HeadCoachID *string  `json:"headCoachId"`
	Name        string   `json:"name"`
	PlaceID     string   `json:"placeId"`
	Published   bool     `json:"published"`
	Writable    bool     `json:"writable"`
}

type TeamInputWithID struct {
	ID    string     `json:"id"`
	Input *TeamInput `json:"input"`
}

type TeamPayload struct {
	Record   *Team  `json:"record"`
	RecordID string `json:"recordId"`
}

type Token struct {
	AccessToken  string    `json:"accessToken"`
	Expiration   time.Time `json:"expiration"`
	RefreshToken string    `json:"refreshToken"`
}

type TrainingConnection struct {
	Edges    []*TrainingEdge `json:"edges"`
	PageInfo *PageInfo       `json:"pageInfo"`
}

type TrainingDayConnection struct {
	Edges    []*TrainingDayEdge `json:"edges"`
	PageInfo *PageInfo          `json:"pageInfo"`
}

type TrainingDayEdge struct {
	Cursor string       `json:"cursor"`
	Node   *TrainingDay `json:"node"`
}

type TrainingDayFilter struct {
	Name *string `json:"name"`
}

type TrainingDayInput struct {
	Day       *DayOfWeek `json:"day"`
	Published bool       `json:"published"`
	StadiumID *string    `json:"stadiumId"`
	TeamID    string     `json:"teamId"`
	Time      *time.Time `json:"time"`
}

type TrainingDayInputWithID struct {
	ID    string            `json:"id"`
	Input *TrainingDayInput `json:"input"`
}

type TrainingDayPayload struct {
	Record   *TrainingDay `json:"record"`
	RecordID string       `json:"recordId"`
}

type TrainingDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type TrainingEdge struct {
	Cursor string    `json:"cursor"`
	Node   *Training `json:"node"`
}

type TrainingFilter struct {
	Name *string `json:"name"`
}

type TrainingInput struct {
	CoachIds    []string  `json:"coachIds"`
	HeadCoachID *string   `json:"headCoachId"`
	Published   bool      `json:"published"`
	StadiumID   *string   `json:"stadiumId"`
	TeamID      string    `json:"teamId"`
	Time        time.Time `json:"time"`
	Visits      int       `json:"visits"`
}

type TrainingInputWithID struct {
	ID    string         `json:"id"`
	Input *TrainingInput `json:"input"`
}

type TrainingPayload struct {
	Record   *Training `json:"record"`
	RecordID string    `json:"recordId"`
}

type UserConnection struct {
	Edges    []*UserEdge `json:"edges"`
	PageInfo *PageInfo   `json:"pageInfo"`
}

type UserDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
}

type UserFilter struct {
	Name *string `json:"name"`
}

type UserInput struct {
	Name      string  `json:"name"`
	Password  *string `json:"password"`
	Phone     string  `json:"phone"`
	Published bool    `json:"published"`
	Role      Role    `json:"role"`
}

type UserInputWithID struct {
	ID    string     `json:"id"`
	Input *UserInput `json:"input"`
}

type UserPayload struct {
	Record   *User  `json:"record"`
	RecordID string `json:"recordId"`
}

type Age string

const (
	AgePreschool Age = "PRESCHOOL"
	AgeSchool    Age = "SCHOOL"
)

var AllAge = []Age{
	AgePreschool,
	AgeSchool,
}

func (e Age) IsValid() bool {
	switch e {
	case AgePreschool, AgeSchool:
		return true
	}
	return false
}

func (e Age) String() string {
	return string(e)
}

func (e *Age) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Age(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Age", str)
	}
	return nil
}

func (e Age) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CoachPaymentRule string

const (
	CoachPaymentRuleByTeam     CoachPaymentRule = "BY_TEAM"
	CoachPaymentRuleByTraining CoachPaymentRule = "BY_TRAINING"
)

var AllCoachPaymentRule = []CoachPaymentRule{
	CoachPaymentRuleByTeam,
	CoachPaymentRuleByTraining,
}

func (e CoachPaymentRule) IsValid() bool {
	switch e {
	case CoachPaymentRuleByTeam, CoachPaymentRuleByTraining:
		return true
	}
	return false
}

func (e CoachPaymentRule) String() string {
	return string(e)
}

func (e *CoachPaymentRule) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CoachPaymentRule(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CoachPaymentRule", str)
	}
	return nil
}

func (e CoachPaymentRule) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "FRIDAY"
	DayOfWeekMonday    DayOfWeek = "MONDAY"
	DayOfWeekSaturday  DayOfWeek = "SATURDAY"
	DayOfWeekSunday    DayOfWeek = "SUNDAY"
	DayOfWeekThursday  DayOfWeek = "THURSDAY"
	DayOfWeekTuesday   DayOfWeek = "TUESDAY"
	DayOfWeekWednesday DayOfWeek = "WEDNESDAY"
)

var AllDayOfWeek = []DayOfWeek{
	DayOfWeekFriday,
	DayOfWeekMonday,
	DayOfWeekSaturday,
	DayOfWeekSunday,
	DayOfWeekThursday,
	DayOfWeekTuesday,
	DayOfWeekWednesday,
}

func (e DayOfWeek) IsValid() bool {
	switch e {
	case DayOfWeekFriday, DayOfWeekMonday, DayOfWeekSaturday, DayOfWeekSunday, DayOfWeekThursday, DayOfWeekTuesday, DayOfWeekWednesday:
		return true
	}
	return false
}

func (e DayOfWeek) String() string {
	return string(e)
}

func (e *DayOfWeek) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DayOfWeek(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DayOfWeek", str)
	}
	return nil
}

func (e DayOfWeek) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Department string

const (
	DepartmentCoach Department = "COACH"
	DepartmentHead  Department = "HEAD"
)

var AllDepartment = []Department{
	DepartmentCoach,
	DepartmentHead,
}

func (e Department) IsValid() bool {
	switch e {
	case DepartmentCoach, DepartmentHead:
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

type LeadSource string

const (
	LeadSourceBrother      LeadSource = "BROTHER"
	LeadSourceFriends      LeadSource = "FRIENDS"
	LeadSourceGis          LeadSource = "GIS"
	LeadSourceGoogle       LeadSource = "GOOGLE"
	LeadSourceInstagram    LeadSource = "INSTAGRAM"
	LeadSourceLift         LeadSource = "LIFT"
	LeadSourceOtherParents LeadSource = "OTHER_PARENTS"
	LeadSourcePaper        LeadSource = "PAPER"
	LeadSourceReturned     LeadSource = "RETURNED"
	LeadSourceTelegram     LeadSource = "TELEGRAM"
	LeadSourceVkontakte    LeadSource = "VKONTAKTE"
	LeadSourceYandex       LeadSource = "YANDEX"
	LeadSourceYandexMap    LeadSource = "YANDEX_MAP"
)

var AllLeadSource = []LeadSource{
	LeadSourceBrother,
	LeadSourceFriends,
	LeadSourceGis,
	LeadSourceGoogle,
	LeadSourceInstagram,
	LeadSourceLift,
	LeadSourceOtherParents,
	LeadSourcePaper,
	LeadSourceReturned,
	LeadSourceTelegram,
	LeadSourceVkontakte,
	LeadSourceYandex,
	LeadSourceYandexMap,
}

func (e LeadSource) IsValid() bool {
	switch e {
	case LeadSourceBrother, LeadSourceFriends, LeadSourceGis, LeadSourceGoogle, LeadSourceInstagram, LeadSourceLift, LeadSourceOtherParents, LeadSourcePaper, LeadSourceReturned, LeadSourceTelegram, LeadSourceVkontakte, LeadSourceYandex, LeadSourceYandexMap:
		return true
	}
	return false
}

func (e LeadSource) String() string {
	return string(e)
}

func (e *LeadSource) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeadSource(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeadSource", str)
	}
	return nil
}

func (e LeadSource) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LeadStatus string

const (
	LeadStatusApproveFirstTraining            LeadStatus = "APPROVE_FIRST_TRAINING"
	LeadStatusCancelAfterApproveFirstTraining LeadStatus = "CANCEL_AFTER_APPROVE_FIRST_TRAINING"
	LeadStatusCancelAfterFirstTraining        LeadStatus = "CANCEL_AFTER_FIRST_TRAINING"
	LeadStatusCancelAfterWrite                LeadStatus = "CANCEL_AFTER_WRITE"
	LeadStatusContracted                      LeadStatus = "CONTRACTED"
	LeadStatusFirstTraining                   LeadStatus = "FIRST_TRAINING"
	LeadStatusFirstTrainingAdd                LeadStatus = "FIRST_TRAINING_ADD"
	LeadStatusWantContract                    LeadStatus = "WANT_CONTRACT"
	LeadStatusWrite                           LeadStatus = "WRITE"
)

var AllLeadStatus = []LeadStatus{
	LeadStatusApproveFirstTraining,
	LeadStatusCancelAfterApproveFirstTraining,
	LeadStatusCancelAfterFirstTraining,
	LeadStatusCancelAfterWrite,
	LeadStatusContracted,
	LeadStatusFirstTraining,
	LeadStatusFirstTrainingAdd,
	LeadStatusWantContract,
	LeadStatusWrite,
}

func (e LeadStatus) IsValid() bool {
	switch e {
	case LeadStatusApproveFirstTraining, LeadStatusCancelAfterApproveFirstTraining, LeadStatusCancelAfterFirstTraining, LeadStatusCancelAfterWrite, LeadStatusContracted, LeadStatusFirstTraining, LeadStatusFirstTrainingAdd, LeadStatusWantContract, LeadStatusWrite:
		return true
	}
	return false
}

func (e LeadStatus) String() string {
	return string(e)
}

func (e *LeadStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LeadStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LeadStatus", str)
	}
	return nil
}

func (e LeadStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MoneyForm string

const (
	MoneyFormBankTransfer MoneyForm = "BANK_TRANSFER"
	MoneyFormCardTransfer MoneyForm = "CARD_TRANSFER"
	MoneyFormCash         MoneyForm = "CASH"
)

var AllMoneyForm = []MoneyForm{
	MoneyFormBankTransfer,
	MoneyFormCardTransfer,
	MoneyFormCash,
}

func (e MoneyForm) IsValid() bool {
	switch e {
	case MoneyFormBankTransfer, MoneyFormCardTransfer, MoneyFormCash:
		return true
	}
	return false
}

func (e MoneyForm) String() string {
	return string(e)
}

func (e *MoneyForm) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MoneyForm(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MoneyForm", str)
	}
	return nil
}

func (e MoneyForm) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderStatus string

const (
	OrderStatusCanceled   OrderStatus = "CANCELED"
	OrderStatusCreated    OrderStatus = "CREATED"
	OrderStatusFinished   OrderStatus = "FINISHED"
	OrderStatusInDelivery OrderStatus = "IN_DELIVERY"
	OrderStatusPayed      OrderStatus = "PAYED"
	OrderStatusReadyToGet OrderStatus = "READY_TO_GET"
)

var AllOrderStatus = []OrderStatus{
	OrderStatusCanceled,
	OrderStatusCreated,
	OrderStatusFinished,
	OrderStatusInDelivery,
	OrderStatusPayed,
	OrderStatusReadyToGet,
}

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusCanceled, OrderStatusCreated, OrderStatusFinished, OrderStatusInDelivery, OrderStatusPayed, OrderStatusReadyToGet:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Priority string

const (
	PriorityHigh    Priority = "HIGH"
	PriorityHighest Priority = "HIGHEST"
	PriorityLow     Priority = "LOW"
	PriorityMiddle  Priority = "MIDDLE"
)

var AllPriority = []Priority{
	PriorityHigh,
	PriorityHighest,
	PriorityLow,
	PriorityMiddle,
}

func (e Priority) IsValid() bool {
	switch e {
	case PriorityHigh, PriorityHighest, PriorityLow, PriorityMiddle:
		return true
	}
	return false
}

func (e Priority) String() string {
	return string(e)
}

func (e *Priority) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Priority(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Priority", str)
	}
	return nil
}

func (e Priority) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleCoach     Role = "COACH"
	RoleDirector  Role = "DIRECTOR"
	RoleEconomist Role = "ECONOMIST"
	RoleEditor    Role = "EDITOR"
	RoleUser      Role = "USER"
)

var AllRole = []Role{
	RoleAdmin,
	RoleCoach,
	RoleDirector,
	RoleEconomist,
	RoleEditor,
	RoleUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleCoach, RoleDirector, RoleEconomist, RoleEditor, RoleUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TaskStatus string

const (
	TaskStatusCancel TaskStatus = "CANCEL"
	TaskStatusDone   TaskStatus = "DONE"
	TaskStatusNew    TaskStatus = "NEW"
	TaskStatusWork   TaskStatus = "WORK"
)

var AllTaskStatus = []TaskStatus{
	TaskStatusCancel,
	TaskStatusDone,
	TaskStatusNew,
	TaskStatusWork,
}

func (e TaskStatus) IsValid() bool {
	switch e {
	case TaskStatusCancel, TaskStatusDone, TaskStatusNew, TaskStatusWork:
		return true
	}
	return false
}

func (e TaskStatus) String() string {
	return string(e)
}

func (e *TaskStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskStatus", str)
	}
	return nil
}

func (e TaskStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type VisitStatus string

const (
	VisitStatusAbsent  VisitStatus = "ABSENT"
	VisitStatusHoliday VisitStatus = "HOLIDAY"
	VisitStatusIll     VisitStatus = "ILL"
	VisitStatusVisited VisitStatus = "VISITED"
)

var AllVisitStatus = []VisitStatus{
	VisitStatusAbsent,
	VisitStatusHoliday,
	VisitStatusIll,
	VisitStatusVisited,
}

func (e VisitStatus) IsValid() bool {
	switch e {
	case VisitStatusAbsent, VisitStatusHoliday, VisitStatusIll, VisitStatusVisited:
		return true
	}
	return false
}

func (e VisitStatus) String() string {
	return string(e)
}

func (e *VisitStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = VisitStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid VisitStatus", str)
	}
	return nil
}

func (e VisitStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
