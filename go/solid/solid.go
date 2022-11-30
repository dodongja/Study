package main

/*
type Report interface {
	Report() string
}

type FinaceReport struct {
	report string
}

func (r *FinaceReport) Report() string {
	return r.report
}

func ReportSender struct{
	...
}

func (s *ReportSender) SendReport(report Report){
	...
}*/

//OCP
//wrong
func SendReport(r *Report, method SendType, receiver string) {
	switch method {
	case Email:
			sendEmail()
	case Fax:
			sendFax()
	case PDF:
			createPDF()
	...
	}
}
//correct
type ReportSender interface {
	Send(r *Report)
}

type EmailSender struct {
}

func (e *EmailSender) Send(r *Report){
	//이메일 전송
}

type FaxSender struct {
}

fucn (f *FaxSender) Send(r *Report){
	//팩스전송
}

//리스코프 치환원칙
type T interface{
	Something()
}

type S struct {
}

func(s *S) Something(){
}

type U struct{	
}

func(u *U) Something(){
}

funcq(t T){
	...
}

var v = &S{}
var u = &U{}

//인터페이스 분리 원칙

//wrong
type Report interface {
	Report() string
	Pages() int
	Author() string
	WrittenDate() time.Time
}

func SendReport(r Report) {
	send(r.Report())
}

//correct
type Report interface {
	Report() string
}

type WrittenInfo interface {
	Pages() int
	Author() string
	WrittenDate() tiem.Time
}

func SendReport(r Report) {
	send(r.Report())
}

//인터페이스 분리 원칙
//wrong
type Mail struct {
	alarm Alarm
}

type Alarm stuct {}

func (m *Mail) OnRecv() {
	m.alarm.Alarm()
}

//correct
type Event interface{
	Register(EventListner)
}

type EventListner interface{
	OnFire()
}

type Mail struct {
	listener EventListner
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.Listener
}

type Alarm struct{

}

func (a *Alarm) OnFire(){
	fmt.Println("알람")
}

var mail = &Mail{}
var listener EventListner = &Alarm{}

mail.Register(listener)
mail.OnRecv()