package patterns

// Observer

type Subscriber interface {
	ReactToPublisherMsg(msg string)
	GetId() int
}

type publisher struct {
	subscriberList []Subscriber
}

func GetNewPublisher() publisher {
	return publisher {subscriberList: make([]Subscriber, 0)}
}

func (pub *publisher) Register(sub Subscriber) {
	pub.subscriberList = append(pub.subscriberList, sub)
}

func (pub *publisher) NotifyAll(msg string){
	for _, sub := range pub.subscriberList {
		println("publisher notifying subscriber id:", sub.GetId(), "named", sub.(subscriber).name)
		sub.ReactToPublisherMsg(msg)
	}
}

type subscriber struct {
	id int
	name string
}

func GetNewSubscriber(id int, name string) subscriber{
	return subscriber{id: id, name: name}
}

func (s subscriber) ReactToPublisherMsg(msg string){
	println(s.name, "reacts to", msg)
}

func (s subscriber) GetId() int {
	return s.id
}


func Observer() {
	pub := GetNewPublisher()
	sub := GetNewSubscriber(1,"first")
	pub.Register(sub)
	pub.NotifyAll("the first message")

	sub2 := GetNewSubscriber(2,"second")
	pub.Register(sub2)
	pub.NotifyAll("the second message")
}
