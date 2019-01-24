// var (
//   notes []string = []string{
//   "A",
//   "A#",
//   "B",
//   "C",
//   "C#",
//   "D",
//   "D#",
//   "E",
//   "F",
//   "F#",
//   "G",
//   "G#",
//   }
// )

// const (
//   keys = 88
// )

// func WhichNote(keyPressCount int) string {
//   idx := ((keyPressCount - 1)%keys)%len(notes)

//   return notes[idx]
// }

// if (m > n) {
//         m, n = n, m <<<<<-- can replace min max function
//     }
// func main() {
// 	var ret [][3]string
// 	m := float64(1)
// 	n := float64(200)

// 	mnmax := max(m,n)

// 	for mnmin := min(m,n); mnmin <= mnmax; mnmin++ {
// 		b := (int(mnmin - 2)) % 7
// 		c := (int(mnmin - 1)) % 9
// 		cval := math.RoundToEven((mnmin-2)/9)
// 		bval := math.RoundToEven((mnmin-2)/7)
// 		if b == 0 && c == 0 {
// 			fmt.Println("Tiok liao!: ", mnmin)
// 			ret = append(ret,[][3]string{{"M: " + fmt.Sprint(mnmin),
// 					"B: "+ fmt.Sprint(bval),
// 					"C: " +fmt.Sprint(cval)}}...)
// 		}
// 	}
// 	fmt.Println(ret)
// }

// func Dist(v, mu float64) float64 {                  // suppose reaction time is 1
//   g := 9.81
//   v = v*(5.0/18.0) //convert to m/s
//   d := (v*v)/(2*mu*g) + v // 1 sec reaction time

//   return d // return in m/s
// }

// func Speed(d, mu float64) float64 {
// 	g := 9.81
// 	b := 2*mu*g
// 	c := -2*d*mu*g
// 	v := (-b + math.Sqrt(b*b - 4*c))/2
// 	v *= 3.6

// 	return v
// }

// func main() {
// 	fmt.Println(Dist(144,0.3))
// 	fmt.Println(Speed(311.83146449201496,0.3))
// }
// func main() {
// 	ip := "192.168.1.300"
// 	s := strings.Split(ip, ".")
// 	if len(s) != 4 {
// 		fmt.Println("false")
// 	}
// 	for _, elem := range s {
// 		i,err := strconv.ParseFloat(elem,64)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		if i/255.0 > 1.0  {
// 			fmt.Println(i/255.0)
// 		}
// 	}
// 	s1 := s[0][0]
// 	fmt.Println(s, string(s1), )
// 	str := "01"
// 	t,err := strconv.Atoi(str)
// 	fmt.Println(t,err)
// }

// func main() {
// 	searchField := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}

// 	for _, searchNumber := range searchField {
// 		fmt.Printf("Searching for %d in list %v\n", searchNumber, searchField)
// 		result, searchCount := binarySearch(searchField, searchNumber)
// 		fmt.Printf("Your number was found in position %d after %d steps.\n\n", result, searchCount)
// 	}
// }

// func binarySearch(a []int, search int) (result int, searchCount int) {
// 	mid := len(a) / 2
// 	switch {
// 	case len(a) == 0:
// 		result = -1 // not found
// 	case a[mid] > search:
// 		result, searchCount = binarySearch(a[:mid], search)
// 	case a[mid] < search:
// 		result, searchCount = binarySearch(a[mid+1:], search)
// 		result += mid + 1
// 	default: // a[mid] == search
// 		result = mid // found
// 	}
// 	searchCount++
// 	return
// }

// func newClosures() (func(), func() int) {
//     a := 0

//     return func() {
//     	a = 5
//     },
//     func() int {
//     	a = a*7
//     	return a
//     }

// }
// func main() {
//     f1, f2 := newClosures()
//     f1() // sets "a" to 5
//     n := f2() // multiplies "a" by 7 - is f2's internal value of "a" 0 or 5 before the call?
//     fmt.Println(n)
// }

// func main() {

// 	saiyans := make(map[string]int)
// 	saiyans["goku"] = 9001
// 	val, exist := saiyans["goku1"]

// 	if exist != true {
// 		fmt.Println(val, exist)
// 	}

// }

// package main

// import (
//     "fmt"

// )

// var rest = []interface{}{}

// func Josephus(items []interface{}, k int) []interface{} {
//   pos := 0
// 	for len(items) > 0 {
// 		pos = (pos + k-1) % len(items)
// 		rest = append(rest, items[pos])
// 		items = append(items[:pos],items[pos+1:]...)
// 	}

// 	return rest
// }

// func main() {

// 	items := []interface{}{1, 2, 3, 4, 5, 6, 7,8,9,10}
// 	rest = Josephus(items, 2)
// 	fmt.Printf("%v, %T", rest,rest)
// }

/*package main

import ("fmt"
import "net/http"
import "strings"
import "strings"
import "io"
import "io"
import "io"
import "io"
import "io"
import "os"
		"math/rand"
		"time")


type walkFn func(*int) walkFn

func pickRandom(fns ...walkFn) walkFn {
	fmt.Println("len",len(fns))
	return fns[rand.Intn(len(fns))]
}

func walkEqual(i *int) walkFn {
	*i += rand.Intn(7) - 3 // -3 -> 3
	return pickRandom(walkForward, walkBackward)
}

func walkForward(i *int) walkFn {
	*i += rand.Intn(6)
	return pickRandom(walkEqual, walkBackward)
}

func walkBackward(i *int) walkFn {
	*i += -rand.Intn(6)
	return pickRandom(walkEqual, walkForward)
}

func main() {
	// time is frozen on playground, so this is actually always
	// the same.  The behavior is different when run locally.
	rand.Seed(time.Now().Unix())

	fn, progress := walkEqual, 0
	for i := 0; i < 20; i++ {
		fn = fn(&progress)
		fmt.Println(progress, )
	}
}
*/

/*package main


import ("fmt"
		"strconv"
		"math/rand"
		"net/http/httptest"
		"net/http"
		"strings"
		"net/url")



type Cart struct {
	items map[string]Item
}

type Item struct {
	ID string
	Name string
	Price float64
	Qty int
}



// AddItem adds an item to the cart
func (c *Cart) AddItem(i Item) {
	c.init()
	ex, ok := c.items[i.ID]
	fmt.Println("in add item: ", ex,ok)
	if existingItem, ok := c.items[i.ID]; ok {
		existingItem.Qty++
		c.items[i.ID] = existingItem
	} else {
		fmt.Println("Got Item")
		i.Qty = 1
		c.items[i.ID] = i

	}
}

func (c *Cart) init() {
	if c.items == nil {
		c.items = map[string]Item{}
	}
}

func main() {
	c := Cart{}
	itemA := Item{ID: "itemA", Name: "Item A", Price: 10.20, Qty: 0}
	itemB := Item{ID: "itemB", Name: "Item B", Price: 7.66, Qty: 0}
	c.AddItem(itemA)
	c.AddItem(itemA)
	c.AddItem(itemB)
	fmt.Println(len(c.items), c.items)
	rand.Seed(4)// Return string rep of arg0
	str := strconv.FormatInt(rand.Int63(), 16)
	w := httptest.NewRecorder()
    params := url.Values{}
    params.Add("username", "user1")
    params.Add("password", "pass1")
	str1 := params.Encode()
	fmt.Println(str,rand.Int63())
	s := strings.NewReader(str1)
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(str1))
	fmt.Println(w, "StatusBadRequest", http.StatusOK, http.StatusTemporaryRedirect)
	fmt.Printf("\ntype %T, and value %v\n", s,s)
	fmt.Println(req,"http unauthorozid:", http.StatusUnauthorized)

}*/

/*package channel

import (
	"sync"
)

var wg sync.WaitGroup

type WG struct {
	main    chan func()
	allDone chan bool
}

func New(n int) WG {
	res := WG{
		main:    make(chan func()),
		allDone: make(chan bool),
	}

	procDone := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for {
				f := <-res.main
				if f == nil { // Close channel, straight away get nil
					procDone <- true
					return
				}
				f()
			}
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			_ = <-procDone
		}
		res.allDone <- true
	}()
	return res
}

func (wg WG) Add(f func()) {
	wg.main <- f
}

func (wg WG) Wait() {
	close(wg.main)
	_ = <-wg.allDone
}
*/
/*package testlab

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ret chan string) {
	defer close(ret)
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ret <- err.Error()
		return
	}

	ret <- fmt.Sprintf("found: %s %q", url, body)

	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, result[i])
	}

	for i := range result {
		for s := range result[i] {
			ret <- s
		}
	}

	return
}

func main() {
	result := make(chan string)
	go Crawl("https://golang.org/", 4, fetcher, result)

	for s := range result {
		fmt.Println("s: ", s)
	}
	fmt.Println(Lenfetch(fetcher, "https://golang.org/"))

}

func Lenfetch(f Fetcher, url string) string {
	_, urls, _ := f.Fetch(url)
	return fmt.Sprintf("%v", len(urls))
}

// fakeFetcher is Fetcher that returns canned results.
//url
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
*/

/*package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what the...")
	}
}

func run() {
	fmt.Printf("running %v as PID %d\n", os.Args[1], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
*/

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // Pool for our struct A
// var pool *sync.Pool

// // A dummy struct with a member
// type A struct {
// 	Name string
// }

// // Func to init pool
// func initPool() {
// 	pool = &sync.Pool{
// 		New: func() interface{} {
// 			fmt.Println("Returning new A")
// 			return new(A)
// 		},
// 	}
// }

// func sad() {
// 	fmt.Println("Sad")
// }

// // Main func
// func main() {
// 	// Initializing pool
// 	initPool()
// 	fmt.Println("Init pool\n")
// 	// Get hold of instance one
// 	one := pool.Get().(*A)

// 	one.Name = "first"
// 	fmt.Printf("one.Name = %s\n", one.Name)
// 	go func(one *A) {
// 		fmt.Printf("Before pool action: %v\n", one.Name)
// 		time.Sleep(5 * time.Second)
// 		fmt.Printf("After Pool kena rompak:%v\n", one.Name)
// 	}(one)
// 	time.Sleep(time.Second)
// 	pool.Put(one)
// 	two := pool.Get().(*A)
// 	two.Name = "two"
// 	fmt.Printf("Two:%v\n", two)
// 	time.Sleep(10 * time.Second)
// }

//
//
//
//

// Mastering Interface :: method with same struct(type) will inherit that interface
// type Shape interface {
// 	GetSides() uint
// 	GetArea() uint
// }

// type Cache struct {
// 	cache map[string]uint
// }

// func (c *Cache) GetValue(str string) int {
// 	value, ok := c.cache[str]
// 	if ok {
// 		fmt.Println("Cache Hit")
// 		return int(value)
// 	} else {
// 		fmt.Println("cache miss, return -1")
// 		return -1
// 	}
// }

// func (c *Cache) SetValue(str string, value int) {
// 	c.cache[str] = uint(value)
// }

// type Square struct {
// 	Cache
// 	Side uint
// }

// type Rectangle struct {
// 	Cache
// 	Width  uint
// 	Height uint
// }

// func (s *Square) GetSides() uint {
// 	value := s.GetValue("Sides")
// 	if value == -1 {
// 		value = int(s.Side * 4)
// 		s.SetValue("Sides", value)
// 	}
// 	return uint(value)
// }

// func (s *Square) GetArea() uint {
// 	value := s.GetValue("Area")
// 	if value == -1 {
// 		value = int(s.Side * s.Side)
// 		s.SetValue("Area", value)

// 	}
// 	return uint(value)
// }

// func (r *Rectangle) GetSides() uint {
// 	value := r.GetValue("Sides")
// 	if value == -1 {
// 		value = int((r.Width + r.Height) * 2)
// 		r.SetValue("Sides", value)

// 	}
// 	return uint(value)

// }
// func (r *Rectangle) GetArea() uint {
// 	value := r.GetValue("Area")
// 	if value == -1 {
// 		value = int(r.Width * r.Height)
// 		r.SetValue("Area", value)

// 	}
// 	return uint(value)
// }
// func (r *Rectangle) Sad() {
// 	fmt.Println("Just print Sad")
// }

// func main() {
// 	shape := []Shape{&Square{Cache{make(map[string]uint)}, 2},
// 		&Rectangle{Cache{make(map[string]uint)}, 2, 3}}

// 	Rect := &Rectangle{Cache{make(map[string]uint)}, 4, 5}
// 	Rect.Sad()

// 	for i := 0; i < 2; i++ {
// 		for _, s := range shape {
// 			fmt.Printf("Sides:%v Area:%v\n", s.GetSides(), s.GetArea())
// 		}
// 	}

// }
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// )

// const (
// 	Sunday  Day = 1 << iota // Sunday == 0
// 	Monday                  // Monday == 1
// 	Tuesday                 // Tuesday == 2
// 	Wednesday
// )

// type HandlerFunc func(string, *request)

// type Handler interface {
// 	ServeHTTP(string, string)
// }

// func newrequest() *request {
// 	return &request{"sad", 2}
// }

// type request struct {
// 	s string
// 	i int
// }

// func (f HandlerFunc) ServeHTTP(w string, r *request) {
// 	fmt.Println(r.s, r.i)
// }

// type functype struct {
// 	numb int
// 	sss  int
// }

// func Afunc(input ...int) int {
// 	var total int
// 	for _, i := range input {
// 		total += i
// 	}
// 	return total
// }

// type Day int

// func fn(in interface{}) {
// 	switch "a" {
// 	case "a":
// 		fmt.Printf("the input is int")
// 	case "b":
// 		fmt.Printf("Input is string")
// 	}
// }

// func Chanwait(c chan int) {
// 	fmt.Println("In Channel wait")
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("i:", i)
// 		c <- i
// 	}
// 	fmt.Println("before close")
// 	close(c)
// }

// type MyType struct {
// 	field string
// }

// type omit *struct{}

// type User struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func NewUser() (*User, error) {
// 	u := &User{}
// 	return u, json.Unmarshal([]byte(`{
// 		"email":"herodotus94@gmail.com",
// 		"password":"1234"
// 	}`), u)
// }

// type BlogPost struct {
// 	URL   string `json:"url"`
// 	Title string `json:"title"`
// }

// type Analytics struct {
// 	Visitors  int `json:"visitors"`
// 	PageViews int `json:"pageviews"`
// }

// func NewBlogPost() (*BlogPost, error) {
// 	bp := &BlogPost{}
// 	return bp, json.Unmarshal([]byte(`{
// 		"url":"www.mrsad.com/sadman",
// 		"title":"Sadman"
// 	}`), bp)
// }

// func NewAnalytics() (*Analytics, error) {
// 	an := &Analytics{}
// 	return an, json.Unmarshal([]byte(`{
// 		"visitors":3333333,
// 		"pageviews":222222
// 	}`), an)
// }

// type Value interface{}

// type CacheItem struct {
// 	key    string `json:"key"`
// 	MaxAge int    `json:"cacheAge"`
// 	Value  Value  `json:"cacheValue"`
// }

// func NewCacheItem() (*CacheItem, error) {
// 	item := &CacheItem{}
// 	return item, json.Unmarshal([]byte(`{
// 		"key":"s2d34ff",
// 		"cacheAge":23,
// 		"cacheValue": {
// 			"nested":true
// 		}
// 	}`), item)
// }
// func main() {
// 	u, _ := NewUser()

// 	json.NewEncoder(os.Stdout).Encode(struct {
// 		*User
// 		Token    string `json:"token"`
// 		Password omit   `json:"password,omitempty"`
// 	}{
// 		User:  u,
// 		Token: "Tolkien",
// 	})
// 	// Composite
// 	bp, _ := NewBlogPost()
// 	an, _ := NewAnalytics()

// 	json.NewEncoder(os.Stdout).Encode(struct {
// 		*BlogPost
// 		*Analytics
// 	}{
// 		bp,
// 		an,
// 	})
// 	//Separate
// 	blogpost, analytic := BlogPost{}, Analytics{}
// 	json.Unmarshal([]byte(`{
// 		"url":"sadman.com/sad",
// 		"title":"sadmaaaan",
// 		"visitors":23323,
// 		"pageviews":555
// 	}`), &struct {
// 		*BlogPost
// 		*Analytics
// 	}{
// 		&blogpost, &analytic,
// 	})

// 	item, _ := NewCacheItem()
// 	json.NewEncoder(os.Stdout).Encode(struct {
// 		*CacheItem
// 		OmitMaxAge omit   `json:"cacheAge,omitempty"`
// 		Omitvalue  omit   `json:"cacheValue,omitempty"`
// 		MaxAge     int    `json:"Max_Age"`
// 		Value      *Value `json:"value"`
// 	}{
// 		CacheItem: item,
// 		MaxAge:    item.MaxAge,
// 		Value:     &item.Value,
// 	})
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// )

// type Record struct {
// 	Author Author `json:"author"`
// 	Title  string `json:"title"`
// 	URL    string `json:"url"`
// }

// type Author struct {
// 	ID    uint64 `json:"id"`
// 	Email string `json:"email"`
// }

// type author Author

// func (a *Author) UnmarshalJSON(b []byte) (err error) {
// 	j, s, n := author{}, "", uint64(0)
// 	if err = json.Unmarshal(b, &j); err == nil {
// 		*a = Author(j)
// 		return
// 	}
// 	if err = json.Unmarshal(b, &s); err == nil {
// 		a.Email = s
// 		return
// 	}
// 	if err = json.Unmarshal(b, &n); err == nil {
// 		a.ID = n
// 	}
// 	return
// }

// func Decode(r io.Reader) (x *Record, err error) {
// 	x = new(Record)
// 	err = json.NewDecoder(r).Decode(x)
// 	return
// }

// func main() {
// 	var r []Record
// 	fmt.Println("ERROR:", json.Unmarshal([]byte(`[{
// 	  "author": "attila@attilaolah.eu",
// 	  "title":  "My Blog",
// 	  "url":    "http://attilaolah.eu"
// 	}, {
// 	  "author": 1234567890,
// 	  "title":  "Westartup",
// 	  "url":    "http://www.westartup.eu"
// 	}, {
// 	  "author": {
// 	    "id":    1234567890,
// 	    "email": "nospam@westartup.eu"
// 	  },
// 	  "title":  "Westartup",
// 	  "url":    "http://www.westartup.eu"
// 	}]`), &r))
// 	fmt.Println("RECORDS:", r)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// type omit *struct{}

// type jsonobject struct {
// 	Title string `json:"title"`
// 	Link  string `json:"link"`
// 	Video *Clip
// }

// type Clip struct {
// 	Video string `json:"video"`
// 	Sound string `json:"sound"`
// }

// func main() {
// 	dajson := &jsonobject{}
// 	daclip := &Clip{
// 		Video: "MrSAAAD",
// 		Sound: "No Sound",
// 	}
// 	b, _ := json.Marshal(daclip)
// 	json.Unmarshal([]byte(`{
// 		"title":"Sad",
// 		"link":"woeopkwokp234241341",
// 		"mp3":1234
// 	}`), dajson)
// 	json.NewEncoder(os.Stdout).Encode(struct {
// 		*jsonobject
// 		*Clip
// 	}{
// 		jsonobject: dajson,
// 		Clip:       daclip,
// 	})
// 	fmt.Printf("\n%s", b)

// 	aslice := []string{"z", "e", "m", "S", "A", "WER"}
// 	sort.Slice(aslice, func(i, j int) bool {
// 		return aslice[j] < aslice[i]
// 	})

// 	fmt.Println(aslice)

// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"net"
// 	"net/http"
// 	"time"
// )

// func Index(w http.ResponseWriter, r *http.Request) {
// 	url := "https://r5---sn-uh-30ad.googlevideo.com/videoplayback?mime=video%2Fwebm&clen=8012260&expire=1538741990&fvip=2&ipbits=0&requiressl=yes&dur=236.111&ms=au%2Crdu&mv=m&mt=1538720269&itag=243&mn=sn-uh-30ad%2Csn-30a7rn7z&key=yt6&gir=yes&aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278&id=o-AKZa0-hwg1KyfnDZ63VMTFX0FL_AeiOF3E8JTdmdLdyz&mm=31%2C29&lmt=1496936483133505&ip=219.92.163.226&ei=hQK3W8-SNIXaowOelpPIAw&pl=18&initcwndbps=508750&sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire&source=youtube&keepalive=yes&c=WEB&alr=yes&signature=03D68FAFFE96EB2FF1771E3CA81BB3EE628193D2.9C66AB47592A6C018DA3526D2F28A2076C10F167&cpn=K4FjVqgAZzSlyIj5&cver=2.20181003&rn=0"

// 	timeout := time.Duration(5) * time.Second
// 	transport := &http.Transport{
// 		ResponseHeaderTimeout: timeout,
// 		Dial: func(network, addr string) (net.Conn, error) {
// 			return net.DialTimeout(network, addr, timeout)
// 		},
// 		DisableKeepAlives: true,
// 	}
// 	client := &http.Client{
// 		Transport: transport,
// 	}
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()

// 	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.
// 	w.Header().Set("Content-Disposition", "attachment; filename=Missing Piece.mp4`")
// 	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
// 	w.Header().Set("Content-Length", r.Header.Get("Content-Length"))

// 	//stream the body to the client without fully loading it into memory
// 	io.Copy(w, resp.Body)
// }

// func main() {
// 	http.HandleFunc("/", Index)
// 	err := http.ListenAndServe(":8000", nil)

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// )

// type User struct {
// 	Id      string
// 	Balance uint64
// }

// func main() {
// 	u := User{Id: "US123", Balance: 8}
// 	b := new(bytes.Buffer)
// 	json.NewEncoder(b).Encode(u)
// 	res, _ := http.Post("https://httpbin.org/post", "application/json; charset=utf-8", b)
// 	var body struct {
// 		// httpbin.org sends back key/value pairs, no map[string][]string
// 		Headers map[string]string `json:"headers"`
// 		Origin  string            `json:"origin"`
// 	}
// 	json.NewDecoder(res.Body).Decode(&body)
// 	fmt.Println(body)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type Vid struct {
// 	Title string   `json:"titlsse"`
// 	Id    []string `json:"ids"`
// }

// const jsonData = `
// 	{"Title": "Sad", "id": "12345"}
// 	{"Title": "wtf", "id": "42342"}
// `

// func main() {

// 	var v *Vid
// 	dapi := new(Vid)
// 	keys := []byte(`{
// 		"titlsse": "Sad",
// 		"ids": ["1234","asds"]
// 		}`)
// 	payload := v.Payload(keys)

// 	fmt.Println(payload)
// 	fmt.Printf("%T, %T", dapi, *dapi)
// }

// func (video *Vid) Payload(key []byte) *Vid {

// 	if err := json.Unmarshal(key, &video); err != nil {
// 		log.Fatalf("Decode err:%v", err)
// 	}

// 	return video
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"time"
// )

// var urls = []string{
// 	"http://pulsoconf.co/",
// 	"http://golang.org/",
// 	"http://matt.aimonetti.net/",
// 	"https://dog.ceo/api/breeds/image/random",
// 	"https://dog.ceo/api/breeds/image/random",
// 	"https://dog.ceo/api/breeds/image/random",
// 	"https://dog.ceo/api/breeds/image/random",
// }

// type HttpResponse struct {
// 	url      string
// 	response *http.Response
// 	err      error
// }

// func asyncHttpGets(urls []string) <-chan *HttpResponse {

// 	ch := make(chan *HttpResponse, len(urls)) // buffered
// 	for _, url := range urls {
// 		go func(url string) {
// 			fmt.Printf("Fetching %s \n", url)
// 			resp, err := http.Get(url)
// 			resp.Body.Close()
// 			ch <- &HttpResponse{url, resp, err}
// 		}(url)
// 	}
// 	return ch
// }

// func HttpGets(url string) *HttpResponse {
// 	fmt.Printf("Fetching %s \n", url)
// 	resp, err := http.Get(url)

// 	defer resp.Body.Close()
// 	return &HttpResponse{url, resp, err}
// }
// func main() {
// 	results := asyncHttpGets(urls)
// 	start := time.Now()
// 	for _ = range urls {
// 		result := <-results
// 		fmt.Printf("%s status: %s\n", result.url, result.response.Status)
// 	}
// 	p := url.Values{"msg": {"saa  d"}}
// 	fmt.Println("haha" + p.Encode())
// 	secs := time.Since(start).Seconds()

// 	start1 := time.Now()
// 	for _, url := range urls {
// 		results := HttpGets(url)
// 		fmt.Printf("%s status: %s\n", results.url, results.response.Status)
// 	}
// 	secs1 := time.Since(start1).Seconds()

// 	fmt.Printf("\n\n Secs: %v    Secs1: %v", secs, secs1)
// }

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"os/exec"
// )

// func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
// 	var out []byte
// 	buf := make([]byte, 1024, 1024)
// 	for {
// 		n, err := r.Read(buf[:])
// 		if n > 0 {
// 			d := buf[:n]
// 			out = append(out, d...)
// 			_, err := w.Write(d)
// 			if err != nil {
// 				return out, err
// 			}
// 		}
// 		if err != nil {
// 			// Read returns io.EOF at the end of file, which is not an error for us
// 			if err == io.EOF {
// 				err = nil
// 			}
// 			return out, err
// 		}
// 	}
// 	// never reached
// 	panic(true)
// 	return nil, nil
// }

// func main() {
// 	path, err := exec.LookPath("go-youtube-dl.exe")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(path)
// 	cmd := exec.Command(path, "--audio-only", "iwyXbD1Rn7g")
// 	// var stdout, stderr []byte
// 	// var errStdout, errStderr error
// 	// stdoutIn, _ := cmd.StdoutPipe()
// 	// stderrIn, _ := cmd.StderrPipe()
// 	// file, _ := os.Create("./sound")
// 	// defer file.Close()
// 	os.Stdout = cmd.Stdout
// 	cmd.Start()

// 	// go func() {
// 	// 	stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
// 	// }()

// 	// go func() {
// 	// 	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
// 	// }()

// 	err = cmd.Wait()
// 	if err != nil {
// 		log.Fatalf("cmd.Run() failed with %s\n", err)
// 	}
// 	// if errStdout != nil || errStderr != nil {
// 	// 	log.Fatalf("failed to capture stdout or stderr\n")
// 	// }
// 	fmt.Println(cmd.Stderr)
// 	// outStr, errStr := string(stdout), string(stderr)
// 	// fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// const (
// 	str  = `[{"playabilityStatus":{"status":"OK"},"streamingData":{"expiresInSeconds":"21540","formats":[{"itag":17,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542219721171327\u0026clen=1024884\u0026ip=115.132.46.180\u0026signature=46ABDC34F990241564E0BDE61750F0C8AA6432CA.65E3CF105EE4861ACA832E9F2A9906B2667367BF\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2F3gpp\u0026gir=yes\u0026dur=105.929\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=17\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026fvip=3\u0026source=youtube\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5531432\u0026key=yt6","mimeType":"video/3gpp; codecs=\"mp4v.20.3, mp4a.40.2\"","bitrate":77485,"width":176,"height":144,"lastModified":"1542219721171327","contentLength":"1024884","quality":"small","qualityLabel":"144p","projectionType":"RECTANGULAR","averageBitrate":77401,"audioQuality":"AUDIO_QUALITY_LOW","approxDurationMs":"105929","audioSampleRate":"22050"},{"itag":18,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542219778734617\u0026ratebypass=yes\u0026clen=9061072\u0026ip=115.132.46.180\u0026signature=6DF17C56443065F6FFFBE302126C00598630AF83.591F4782069E30737A0E9A88E6C6B9E443D168CC\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fmp4\u0026gir=yes\u0026dur=105.882\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire\u0026itag=18\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026fvip=3\u0026source=youtube\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5531432\u0026key=yt6","mimeType":"video/mp4; codecs=\"avc1.42001E, mp4a.40.2\"","bitrate":685063,"width":640,"height":360,"lastModified":"1542219778734617","contentLength":"9061072","quality":"medium","qualityLabel":"360p","projectionType":"RECTANGULAR","averageBitrate":684616,"audioQuality":"AUDIO_QUALITY_LOW","approxDurationMs":"105882","audioSampleRate":"44100"},{"itag":22,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?mt=1542270684\u0026itag=22\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026txp=5531432\u0026expire=1542292393\u0026lmt=1542221909361935\u0026ip=115.132.46.180\u0026ratebypass=yes\u0026c=WEB\u0026source=youtube\u0026fvip=3\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026pl=19\u0026mv=m\u0026initcwndbps=512500\u0026ipbits=0\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026key=yt6\u0026signature=177E85825B8628DA078730530DCEFE8343EE0B30.A95DE357DF23F8BAD695970A9F1DD35A6A0AC3AB\u0026mime=video%2Fmp4\u0026dur=105.882\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=dur%2Cei%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire","mimeType":"video/mp4; codecs=\"avc1.64001F, mp4a.40.2\"","bitrate":1510578,"width":1280,"height":720,"lastModified":"1542221909361935","quality":"hd720","qualityLabel":"720p","projectionType":"RECTANGULAR","audioQuality":"AUDIO_QUALITY_MEDIUM","approxDurationMs":"105882","audioSampleRate":"44100"},{"itag":36,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542219720063160\u0026clen=3015152\u0026ip=115.132.46.180\u0026signature=01FB0248BBD672A8E047E1665E38D20D64DC89A5.3B49568EE00392463662D848CFEAF12E356E8DF7\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2F3gpp\u0026gir=yes\u0026dur=105.929\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=36\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026fvip=3\u0026source=youtube\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5531432\u0026key=yt6","mimeType":"video/3gpp; codecs=\"mp4v.20.3, mp4a.40.2\"","bitrate":227958,"width":320,"height":180,"lastModified":"1542219720063160","contentLength":"3015152","quality":"small","qualityLabel":"144p","projectionType":"RECTANGULAR","averageBitrate":227711,"audioQuality":"AUDIO_QUALITY_LOW","approxDurationMs":"105929","audioSampleRate":"22050"},{"itag":43,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542225427955733\u0026ratebypass=yes\u0026clen=11376606\u0026ip=115.132.46.180\u0026signature=C6FE4D85EFC2D166FF47B66CEF60BFFB8126F052.AC5A34F92CE7825FD15B283BBD5E7C56E361360A\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fwebm\u0026gir=yes\u0026dur=0.000\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=clen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Cratebypass%2Crequiressl%2Csource%2Cexpire\u0026itag=43\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026fvip=3\u0026source=youtube\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5511222\u0026key=yt6","mimeType":"video/webm; codecs=\"vp8.0, vorbis\"","bitrate":2147483647,"width":640,"height":360,"lastModified":"1542225427955733","contentLength":"11376606","quality":"medium","qualityLabel":"360p","projectionType":"RECTANGULAR","audioQuality":"AUDIO_QUALITY_MEDIUM"}],"adaptiveFormats":[{"itag":137,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542221454435266\u0026aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278\u0026clen=30436276\u0026ip=115.132.46.180\u0026signature=6AB20A7BD452C4713BF7D73B8904A7A06DF7BCD1.108301514E577FCCA4FC5F2221F72F2B2E9868C3\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fmp4\u0026gir=yes\u0026dur=105.814\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=137\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026keepalive=yes\u0026source=youtube\u0026fvip=3\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5533432\u0026key=yt6","mimeType":"video/mp4; codecs=\"avc1.640028\"","bitrate":4404072,"width":1920,"height":1080,"initRange":{"start":"0","end":"764"},"indexRange":{"start":"765","end":"1060"},"lastModified":"1542221454435266","contentLength":"30436276","quality":"hd1080","fps":24,"qualityLabel":"1080p","projectionType":"RECTANGULAR","averageBitrate":2301115,"approxDurationMs":"105814"},{"itag":248,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542223009095510\u0026aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278\u0026clen=30993462\u0026ip=115.132.46.180\u0026signature=CACD3FA85B2FCCB990D162BFCE1448430ECF0584.254609E382C408AA97EF8CE08155FDB568FFC492\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fwebm\u0026gir=yes\u0026dur=105.814\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=248\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026keepalive=yes\u0026source=youtube\u0026fvip=3\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5530432\u0026key=yt6","mimeType":"video/webm; codecs=\"vp9\"","bitrate":2656073,"width":1920,"height":1080,"initRange":{"start":"0","end":"219"},"indexRange":{"start":"220","end":"594"},"lastModified":"1542223009095510","contentLength":"30993462","quality":"hd1080","fps":24,"qualityLabel":"1080p","projectionType":"RECTANGULAR","averageBitrate":2343240,"colorInfo":{"primaries":"COLOR_PRIMARIES_BT709","transferCharacteristics":"COLOR_TRANSFER_CHARACTERISTICS_BT709","matrixCoefficients":"COLOR_MATRIX_COEFFICIENTS_BT709"},"approxDurationMs":"105814"},{"itag":136,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542221594582036\u0026aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278\u0026clen=18272668\u0026ip=115.132.46.180\u0026signature=564663AAF1F76F6FC69F59EDC58DCE74576F7BF9.4EB6B72CFAA838F7CFC9B5CB87D81344A155D967\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fmp4\u0026gir=yes\u0026dur=105.814\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=136\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026keepalive=yes\u0026source=youtube\u0026fvip=3\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5533432\u0026key=yt6","mimeType":"video/mp4; codecs=\"avc1.4d401f\"","bitrate":2248232,"width":1280,"height":720,"initRange":{"start":"0","end":"762"},"indexRange":{"start":"763","end":"1058"},"lastModified":"1542221594582036","contentLength":"18272668","quality":"hd720","fps":24,"qualityLabel":"720p","projectionType":"RECTANGULAR","averageBitrate":1381493,"approxDurationMs":"105814"},{"itag":247,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542223103090424\u0026aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278\u0026clen=14518530\u0026ip=115.132.46.180\u0026signature=41C13A8323C952D2433FB223E25DE924DAD147B5.A833FE7C39698392EDC323FA76570F4EB82E397D\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fwebm\u0026gir=yes\u0026dur=105.814\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=247\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026keepalive=yes\u0026source=youtube\u0026fvip=3\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5530432\u0026key=yt6","mimeType":"video/webm; codecs=\"vp9\"","bitrate":1517619,"width":1280,"height":720,"initRange":{"start":"0","end":"219"},"indexRange":{"start":"220","end":"584"},"lastModified":"1542223103090424","contentLength":"14518530","quality":"hd720","fps":24,"qualityLabel":"720p","projectionType":"RECTANGULAR","averageBitrate":1097664,"colorInfo":{"primaries":"COLOR_PRIMARIES_BT709","transferCharacteristics":"COLOR_TRANSFER_CHARACTERISTICS_BT709","matrixCoefficients":"COLOR_MATRIX_COEFFICIENTS_BT709"},"approxDurationMs":"105814"},{"itag":135,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542221349672976\u0026aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278\u0026clen=11013263\u0026ip=115.132.46.180\u0026signature=04A689FB1720C9D4E1C2BF0DE0D7DB9D66C77769.4C4C379431734F155E38FCC30996AD37523D290D\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fmp4\u0026gir=yes\u0026dur=105.814\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=135\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026keepalive=yes\u0026source=youtube\u0026fvip=3\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5533432\u0026key=yt6","mimeType":"video/mp4; codecs=\"avc1.4d401e\"","bitrate":1199591,"width":854,"height":480,"initRange":{"start":"0","end":"762"},"indexRange":{"start":"763","end":"1058"},"lastModified":"1542221349672976","contentLength":"11013263","quality":"large","fps":24,"qualityLabel":"480p","projectionType":"RECTANGULAR","averageBitrate":832650,"approxDurationMs":"105814"},{"itag":244,"url":"https://r7---sn-uh-30aek.googlevideo.com/videoplayback?expire=1542292393\u0026lmt=1542222776453587\u0026aitags=133%2C134%2C135%2C136%2C137%2C160%2C242%2C243%2C244%2C247%2C248%2C278\u0026clen=7769312\u0026ip=115.132.46.180\u0026signature=7D9CFE54BB9048AE87AE199055CFE455EFDF1256.6B3A7B748DB3ED387ADFF580B7A84A1A8EE62B78\u0026mv=m\u0026mt=1542270684\u0026ms=au%2Crdu\u0026mm=31%2C29\u0026c=WEB\u0026mn=sn-uh-30aek%2Csn-30a7yn76\u0026mime=video%2Fwebm\u0026gir=yes\u0026dur=105.814\u0026id=o-ANKI0douR-y5nHT8vo0_4c0SMqrRUo5-CcLX_GvGSi8T\u0026sparams=aitags%2Cclen%2Cdur%2Cei%2Cgir%2Cid%2Cinitcwndbps%2Cip%2Cipbits%2Citag%2Ckeepalive%2Clmt%2Cmime%2Cmm%2Cmn%2Cms%2Cmv%2Cpl%2Crequiressl%2Csource%2Cexpire\u0026itag=244\u0026requiressl=yes\u0026ei=SS_tW9bJEYWBowO8irjgDg\u0026keepalive=yes\u0026source=youtube\u0026fvip=3\u0026pl=19\u0026initcwndbps=512500\u0026ipbits=0\u0026txp=5530432\u0026key=yt6","mimeType":"video/webm; codecs=\"vp9\"","bitrate":752162,"width":854,"height":480,"initRange":{"start":"0","end":"219"},"indexRange":{"start":"220","end":"584"},"lastModified":"1542222776453587","contentLength":"7769312","quality":"large","fps":24,"qualityLabel":"480p","projectionType":"RECTANGULAR","averageBitrate":587393,"colorInfo":{"primaries":"COLOR_PRIMARIES_BT709","transferCharacteristics":"COLOR_TRANSFER_CHARACTERISTICS_BT709","matrixCoefficients":"COLOR_MATRIX_COEFFICIENTS_BT709"},"approxDurationMs":"105814"},`
// 	str1 = `https%3A%2F%2Fr7---sn-uh-30aek.googlevideo.com%2Fvideoplayback%3Fmv%3Dm%26mt%3D1542277213%26signature%3D882E1331DB062079F83514E49332024CF65C0EBC.969CAB74C5C0B89C89E0D587CCE0020982D54CCE%26ms%3Dau%252Crdu%26mn%3Dsn-uh-30aek%252Csn-30a7rn7s%26mm%3D31%252C29%26c%3DWEB%26fvip%3D3%26expire%3D1542298916%26key%3Dyt6%26mime%3Dvideo%252Fmp4%26ipbits%3D0%26lmt%3D1542219778734617%26itag%3D18%26dur%3D105.882%26pl%3D19%26initcwndbps%3D551250%26source%3Dyoutube%26ratebypass%3Dyes%26requiressl%3Dyes%26gir%3Dyes%26txp%3D5531432%26id%3Do-AJIZXajxQmUtwR86-VEh7Nf69yTSthVbBKTYWhC_s7aj%26clen%3D9061072%26ip%3D115.132.46.180%26sparams%3Dclen%252Cdur%252Cei%252Cgir%252Cid%252Cinitcwndbps%252Cip%252Cipbits%252Citag%252Clmt%252Cmime%252Cmm%252Cmn%252Cms%252Cmv%252Cpl%252Cratebypass%252Crequiressl%252Csource%252Cexpire%26ei%3DxEjtW63LG4jKowPC-aKIAg`
// )

// func main() {
// 	PResponseList := strings.Split(str, `},{`)
// 	for _, ListItem := range PResponseList {
// 		matched, _ := regexp.MatchString(`"itag":22`, ListItem)
// 		if matched {
// 			fmt.Println(ListItem)
// 			re := regexp.MustCompile(`https:\/\/[^"]+`)
// 			urls := re.FindStringSubmatch(ListItem)
// 			result := strings.Replace(urls[0], `\u0026`, "&", -1)
// 			fmt.Println(result)
// 			break
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/cookiejar"
// 	"net/url"
// 	"strings"
// )

// func main() {
// 	jar, _ := cookiejar.New(nil)

// 	var cookies []*http.Cookie

// 	firstCookie := &http.Cookie{
// 		Name:   "PREF",
// 		Value:  "f1=50000000&f5=30",
// 		Path:   "/",
// 		Domain: ".youtube.com",
// 	}

// 	cookies = append(cookies, firstCookie)

// 	secondCookie := &http.Cookie{
// 		Name:   "VISITOR_INFO1_LIVE",
// 		Value:  "Luy_gz784rQ",
// 		Path:   "/",
// 		Domain: ".youtube.com",
// 	}

// 	cookies = append(cookies, secondCookie)

// 	thirdCookie := &http.Cookie{
// 		Name:   "YSC",
// 		Value:  "9jGyTn_JiBk",
// 		Path:   "/",
// 		Domain: ".youtube.com",
// 	}

// 	cookies = append(cookies, thirdCookie)

// 	fourthCookie := &http.Cookie{
// 		Name:   "dkv",
// 		Value:  "19a17a80fd703efd450d5ef9dadc32cee3QEAAAAdGxpcGn9mTBVMA==",
// 		Path:   "/",
// 		Domain: ".youtube.com",
// 	}

// 	cookies = append(cookies, fourthCookie)

// 	// URL for cookies to remember. i.e reply when encounter this URL
// 	cookieURL, _ := url.Parse("https://www.youtube.com/results?search_query=")

// 	jar.SetCookies(cookieURL, cookies)
// 	fmt.Println("Url", cookieURL)
// 	// sanity check
// 	fmt.Printf("\n%v\n", jar.Cookies(cookieURL))

// 	//setup our client based on the cookies data
// 	client := &http.Client{
// 		Jar: jar,
// 	}

// 	urlData := url.Values{}
// 	urlData.Set("search_query", "macross")

// 	req, _ := http.NewRequest("POST", "https://www.youtube.com/results?search_query=", strings.NewReader(urlData.Encode()))
// 	req.AddCookie(firstCookie)
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	fmt.Println("\n\nCookies sent\n", req.Cookies())
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(nil)
// 	}
// 	fmt.Println("\nReturned Cookies", jar.Cookies(cookieURL))

// 	// body, _ := ioutil.ReadAll(resp.Body)
// 	resp.Body.Close()

// 	// // display content to screen ... save this to a HTML file and view the file with browser ;-)
// 	// fmt.Println(string(body))
// }

// package main

// import (
// 	"net/http"
// )

// func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
// 	html := "Hello World! "
// 	w.Write([]byte(html))
// }

// func ReadCookie(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("ithinkidroppedacookie")
// 	if err != nil {
// 		w.Write([]byte("error in reading cookie : " + err.Error() + "\n"))
// 	} else {
// 		value := c.Value
// 		w.Write([]byte("cookie has : " + value + "\n"))
// 	}
// }

// // see https://golang.org/pkg/net/http/#Cookie
// // Setting MaxAge<0 means delete cookie now.

// func DeleteCookie(w http.ResponseWriter, r *http.Request) {
// 	c := http.Cookie{
// 		Name:   "ithinkidroppedacookie",
// 		MaxAge: -1}
// 	http.SetCookie(w, &c)

// 	w.Write([]byte("old cookie deleted!\n"))
// }

// func CreateCookie(w http.ResponseWriter, r *http.Request) {
// 	c := http.Cookie{
// 		Name:   "ithinkidroppedacookie",
// 		Value:  "thedroppedcookiehasgoldinit",
// 		MaxAge: 3600}
// 	http.SetCookie(w, &c)

// 	w.Write([]byte("new cookie created!\n"))
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", SayHelloWorld)
// 	mux.HandleFunc("/readcookie", ReadCookie)
// 	mux.HandleFunc("/deletecookie", DeleteCookie)
// 	mux.HandleFunc("/createcookie", CreateCookie)
// 	http.ListenAndServe(":8080", mux)
// }

// package main

// import (
// 	"bytes"
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"fmt"
// )

// type ecb struct {
// 	b         cipher.Block
// 	blockSize int
// }

// func newECB(b cipher.Block) *ecb {
// 	return &ecb{
// 		b:         b,
// 		blockSize: b.BlockSize(),
// 	}
// }

// type ecbEncrypter ecb

// func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
// 	return (*ecbEncrypter)(newECB(b))
// }

// func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

// func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
// 	if len(src)%x.blockSize != 0 {
// 		panic("crypto/cipher: input not full blocks")
// 	}
// 	if len(dst) < len(src) {
// 		panic("crypto/cipher: output smaller than input")
// 	}
// 	for len(src) > 0 {
// 		x.b.Encrypt(dst, src[:x.blockSize])
// 		src = src[x.blockSize:]
// 		dst = dst[x.blockSize:]
// 	}
// }

// type ecbDecrypter ecb

// func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
// 	return (*ecbDecrypter)(newECB(b))
// }

// func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

// func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
// 	if len(src)%x.blockSize != 0 {
// 		panic("crypto/cipher: input not full blocks")
// 	}
// 	if len(dst) < len(src) {
// 		panic("crypto/cipher: output smaller than input")
// 	}
// 	for len(src) > 0 {
// 		x.b.Decrypt(dst, src[:x.blockSize])
// 		src = src[x.blockSize:]
// 		dst = dst[x.blockSize:]
// 	}
// }

// func Pad(src []byte) []byte {
// 	padding := aes.BlockSize - len(src)%aes.BlockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(src, padtext...)
// }

// func main() {
// 	key := []byte("jo6aey6haid2Teih")
// 	plaintext := Pad([]byte("abc"))

// 	if len(plaintext)%aes.BlockSize != 0 {
// 		panic("plaintext is not a multiple of block size")
// 	}

// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	ciphertext := make([]byte, len(plaintext))
// 	mode := NewECBEncrypter(block)
// 	mode.CryptBlocks(ciphertext, plaintext)
// 	fmt.Printf("%x\n", ciphertext)
// }

// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"golang.org/x/crypto/blowfish"
// 	"io"
// )

// func main() {
// 	// ¤
// 	position := 0
// 	var destBuffer bytes.Buffer
// 	img := bytes.NewReader([]byte("sdseagojsdjgsgjdsfogdjsj"))
// 	streamLength := img.Len()
// 	for position < streamLength {
// 		dst := make([]byte, 1)
// 		io.ReadFull(img, dst)
// 		m, _ := destBuffer.Write(dst)
// 		fmt.Println(m)
// 		fmt.Println(destBuffer.String())
// 		position += 1
// 	}
// 	var a int64 = 3
// 	fmt.Printf("\n%T %v\n", int(a), int(a))

// 	decrypter, err := blowfish.NewCipher([]byte("llfk9f,7e%u`<d49"))
// 	fmt.Println(decrypter, err)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	fmt.Println("in test")
// 	cmd := exec.Command("testlab.exe")
// 	cmd.Stdout = os.Stdout
// 	err := cmd.Start()
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = cmd.Wait()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )

// func main() {

// 	scanner := bufio.NewScanner(os.Stdin)
// 	for scanner.Scan() {
// 		InputText := scanner.Text()
// 		matchTableIn, _ := regexp.MatchString("[0-9 ]+", InputText)
// 		if matchTableIn {
// 			//Get Basic infos
// 			Str1 := strings.Fields(InputText)
// 			intStr1, _ := strconv.Atoi(Str1[0])
// 			intStr2, _ := strconv.Atoi(Str1[1])
// 			if intStr1 > 500000 {
// 				fmt.Printf("N Out Of range")
// 			}
// 			if intStr2 > 100000 {
// 				fmt.Printf("M Out oF range")
// 			}
// 			var IntSlice = make([]int, intStr1)
// 			var FinalString string

// 			// Inner code
// 			scanInner := bufio.NewScanner(os.Stdin)

// 			for i := 0; i < intStr2; i++ {
// 				for scanInner.Scan() {
// 					InputTextInner := scanInner.Text()
// 					matchNoteIn, _ := regexp.MatchString("[0-9 ]+", InputTextInner)
// 					if matchNoteIn {
// 						Str1 = strings.Fields(InputTextInner)
// 						val1, _ := strconv.Atoi(Str1[0])
// 						val2, _ := strconv.Atoi(Str1[1])
// 						val3, _ := strconv.Atoi(Str1[2])
// 						if val3 > 5 {
// 							fmt.Println("C Out of Range")
// 							os.Exit(0)
// 						}
// 						for val1 <= val2 {
// 							IntSlice[val1-1] += val3
// 							val1++
// 						}
// 					} else {
// 						fmt.Println("Input needs to be: int1 int2 int3")
// 						os.Exit(0)
// 					}
// 					break
// 				}
// 			}
// 			scanLast := bufio.NewScanner(os.Stdin)
// 			fmt.Println("entering last")
// 			for scanLast.Scan() {
// 				InputText := scanLast.Text()
// 				Str1 := strings.Fields(InputText)
// 				intStr1, _ := strconv.Atoi(Str1[0])
// 				intStr2, _ := strconv.Atoi(Str1[1])
// 				IntSlice = IntSlice[intStr1-1 : intStr2]
// 				break
// 			}
// 			for _, val := range IntSlice {
// 				FinalString += fmt.Sprintf("%v ", val)
// 			}
// 			fmt.Println("final Stirng", FinalString)

// 		} else {
// 			fmt.Println("Incorrect Input Format")
// 			break
// 		}
// 		break
// 	}

// 	if scanner.Err() != nil {
// 		fmt.Println("Aww Snap Error reading input")
// 	}
// }

// package testlab

// import (
// 	"fmt"
// 	"strings"
// )

// func StringConstructing(a, s string) int {
// 	var c int
// 	var counter int
// 	for s != "" {
// 		s, counter = TrimMatch(a, s)

// 		fmt.Println("Counter: ", counter, s)
// 		c += counter
// 	}

// 	return c
// }

// func TrimMatch(a, s string) (string, int) {
// 	// Trim Perfect match
// 	var count int
// 	if strings.HasPrefix(s, a) {
// 		fmt.Println("perfect match")
// 		s = strings.TrimPrefix(s, a)
// 		return s, 1
// 	}

// 	for a != "" {
// 		if strings.HasPrefix(s, string(a[0])) {
// 			s = strings.TrimPrefix(s, string(a[0]))
// 		} else {
// 			count++
// 			a = a[1:]
// 		}
// 		fmt.Println("a inside", a)
// 	}
// 	fmt.Println(a)
// 	return s, count
// }

// package testlab

// import (
// 	"fmt"
// 	"math"
// )

// func LastDigit(a []int) int {
// 	if len(a) == 0 {
// 		return 0
// 	}
// var value float64
// var lastValPow float64
// var lastValue float64
// lastValPow = math.Mod(float64(a[len(a)-1]), 100)
// lastValue = math.Mod(float64(a[len(a)-2]), 10)
// 	if lastValPow == 0 && lastValue != 0 {
// 		value = 1
// 	} else {
// 		value = math.Pow(lastValue, math.Mod(lastValPow, 4)+4)
// 	}

// 	if lastValPow == 0 && lastValue == 0 {
// 		value = 1
// 	}
// 	fmt.Println("lastValPow", lastValPow)
// 	fmt.Println("lastVal", lastValue)
// 	fmt.Println("value", value)
// 	for i := len(a) - 3; i >= 0; i-- {
// 		value = math.Mod(value, 100)
// 		lastValue = math.Mod(float64(a[i]), 10)
// 		if value == 1 {
// 			value = lastValue
// 		} else if value == 0 && lastValue == 0 {
// 			value = 0
// 		} else if value == 0 {
// 			value = 1
// 		} else {
// 			value = math.Pow(float64(lastValue), math.Mod(value, 4)+4)
// 		}
// 		fmt.Println(i, value)
// 	}

// 	return int(math.Mod(value, 10))
// }

// import (
// 	"fmt"
// 	"math"
// )

// func LastDigit(a []int) int {
// 	if len(a) == 0 {
// 		return 1
// 	}
// 	var value float64
// 	var lastValPow float64
// 	var lastValue float64

// 	for i := len(a) - 1; i >= 1; i-- {
// 		lastValPow = math.Mod(float64(a[i]), 100)
// 		lastValue = math.Mod(float64(a[i-1]), 10)
// 		if (lastValPow == 0 && lastValue == 0) || lastValPow == 0 {
// 			value = 1
// 		} else if lastValue == 0 {
// 			value = 0
// 		} else if i == len(a)-1 {
// 			value = math.Pow(lastValue, math.Mod(lastValPow, 4)+4)
// 		} else if value == 1 {
// 			value = lastValue
// 		} else {
// 			value = math.Pow(lastValue, math.Mod(value, 4)+4)
// 		}
// 		a[i-1] = int(value)
// 		fmt.Println(lastValPow, lastValue, value)
// 	}
// 	return int(math.Mod(value, 10))
// }
// import (
// 	"fmt"
// 	"math"
// )

// func LastDigit(a []int) int {
// 	if len(a) == 0 {
// 		return 1
// 	}
// 	if len(a) == 1 {
// 		return a[0] % 10
// 	}
// 	b := make([]int, len(a))
// 	copy(b, a)
// 	fmt.Println("a", a)
// 	var value float64
// 	var lastValPow float64
// 	var lastValue float64

// 	for i := len(a) - 1; i >= 1; i-- {
// 		lastValPow = math.Mod(float64(a[i]), 100)
// 		lastValue = math.Mod(float64(a[i-1]), 10)
// 		if (lastValPow == 0 && lastValue == 0) || (lastValPow == 0 && b[i] == 0) {
// 			value = 1
// 		} else if lastValue == 0 {
// 			value = 0
// 		} else if i == len(a)-1 && lastValPow != 1 {
// 			value = math.Pow(lastValue, math.Mod(lastValPow, 4)+4)
// 		} else if value == 1 || lastValPow == 1 {
// 			value = lastValue
// 		} else {
// 			value = math.Pow(lastValue, math.Mod(value, 4)+4)
// 		}
// 		a[i-1] = int(value)
// 		fmt.Println(lastValue, lastValPow, value)
// 	}
// 	return int(math.Mod(value, 10))
// }

// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"sync"
// )

// func main() {
// 	urls := []string{
// 		"http://www.reddit.com/r/aww.json",
// 		"http://www.reddit.com/r/funny.json",
// 	}
// 	jsonResponses := make(chan string)

// 	var wg sync.WaitGroup

// 	wg.Add(len(urls))

// 	for _, url := range urls {
// 		go func(url string) {
// 			defer wg.Done()
// 			res, err := http.Get(url)
// 			if err != nil {
// 				log.Fatal(err)
// 			} else {
// 				defer res.Body.Close()
// 				body, err := ioutil.ReadAll(res.Body)
// 				if err != nil {
// 					log.Fatal(err)
// 				} else {
// 					jsonResponses <- string(body)
// 				}
// 			}
// 		}(url)
// 	}

// 	go func() {
// 		for response := range jsonResponses {
// 			fmt.Println(response)
// 		}
// 	}()

// 	wg.Wait()
// }

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//Test channels
func main() {
	fmt.Println("Start")
	chDone := make(chan string, 100)
	ch := make(chan int)

	go blocker(ch)
	fmt.Println("Ended", <-ch)

	theMine := [5]string{"ore1", "ore2", "ore3"}
	oreChan := make(chan string)
	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			oreChan <- item //send
		}
	}(theMine)
	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChan //receive
			fmt.Println("Miner: Received " + foundOre + " from finder")
		}
		chDone <- "Done!"
	}()
	<-chDone //Done Channel, blocks till receive value

	bufferedChan := make(chan string, 3)
	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()
	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
		chDone <- "Done"
	}()
	<-chDone //Done Channel, blocks till receive value

	theMine = [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	// Finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //send item on oreChannel
			}
		}
	}(theMine)
	// Ore Breaker
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel //read from oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //send to minedOreChan
		}
	}()
	// Smelter
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan //read from minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
			chDone <- "Done"
		}
	}()
	<-chDone //Done Channel, blocks till receive value // Again, you can ignore this

	myChan := make(chan string)

	go func() {
		myChan <- "Message!"
	}()

	select {
	case msg := <-myChan:
		fmt.Println(msg)
		chDone <- "Done"
	default:
		fmt.Println("No Msg")
	}
	<-chDone //Done Channel, blocks till receive value
	select {
	case msg := <-myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(2*time.Second))
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		myfunc(ctx)
	}()

	wg.Wait()
	fmt.Printf("In main, ctx err is %+v\n", ctx.Err())

}

func myfunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Done is called when time out
			fmt.Printf("Ctx is kicking in with error:%+v\n", ctx.Err())
			return
		case <-time.After(3 * time.Second):
			fmt.Printf("I was not canceled\n")
			return
		}
	}
}

func blocker(ch chan int) {
	fmt.Println("blocker")
	time.Sleep(2 * time.Second)
	ch <- 0
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"reflect"
// 	"time"
// )

// type ConfigMap func(map[string]interface{}) DatastoreConfig

// type DatastoreConfig interface {
// 	Something() string
// }

// type MapTypeString map[string]string

// type AnotherData struct {
// 	Number int
// }

// var datastores map[string]ConfigMap

// const Int = 12
// const Int64 int64 = 13

// func main() {
// 	fmt.Println("TypeOf: ", Int*time.Second)
// 	fmt.Println("TypeOf: ", reflect.TypeOf(Int64))
// 	params := make(map[string]interface{})
// 	params["number"] = 12
// 	params["string"] = "sad"
// 	datastore := map[string]ConfigMap{
// 		"sad": SadFunc,
// 	}
// 	which, ok := params["string"].(string)
// 	if !ok {
// 		log.Fatal("not a string")
// 	}
// 	fmt.Println("which", which)

// 	fun, ok := datastore[which]
// 	if !ok {
// 		log.Fatalf("Dont have datastore[%s]", which)
// 	}
// 	pointer := fun(params)
// 	fmt.Println(pointer.Something(), "\n\n\n")

// 	newparams := make(MapTypeString)
// 	newparams["walao"] = "haha"
// 	fmt.Println("newparams", reflect.TypeOf(newparams["wala"]))

// 	err := os.MkdirAll("store", os.ModePerm)
// 	if err != nil {
// 		fmt.Errorf("%v", err)
// 	}
// 	file, err := os.Create("store/sad2.txt")
// 	if err != nil {
// 		fmt.Errorf("fck Error %v", err)
// 	}
// 	fmt.Fprint(file, "ss")
// }

// func SadFunc(params map[string]interface{}) DatastoreConfig {
// 	var a AnotherData
// 	fmt.Println("a", a.Number)
// 	return &a
// }

// func (a *AnotherData) Something() string {
// 	return "from AnotherData"
// }

// func Something() string {
// 	return "hahha"
// }

// IPFS using Plugins

// package main

// import (
// 	"context"
// 	"fmt"
// 	"github.com/ipfs/go-ipfs/core"
// 	"github.com/ipfs/go-ipfs/core/coreapi"
// 	"github.com/ipfs/go-ipfs/core/coreapi/interface"
// 	"github.com/ipfs/go-ipfs/repo/fsrepo"
// 	"io/ioutil"
// 	"os"
// 	"testlab/lib"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
// 	defer cancel()
// 	_, err := lib.LoadPlugins("")
// 	OnError(err)
// 	r, err := fsrepo.Open("C:/Users/SiangHwa/.ipfs")
// 	OnError(err)
// 	cfg := &core.BuildCfg{
// 		Repo: r,
// 	}
// 	node, err := core.NewNode(ctx, cfg)
// 	OnError(err)
// 	path, err := iface.ParsePath("QmVQt8kV1HsNBBWUGYiKNup3j6EaCtBun2bMUcoG6iFCCd")
// 	api, err := coreapi.NewCoreAPI(node)
// 	OnError(err)

// 	list, err := api.Object().Data(node.Context(), path)
// 	OnError(err)

// 	data, err := ioutil.ReadAll(list)
// 	fmt.Println("Data", string(data[:]))

// }

// func OnError(err error) {
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		panic(err)
// 	}
// }

// IPFS using HTTP APi

// package main

// import (
// 	"fmt"
// 	"io/ioutil"

// 	shell "github.com/ipfs/go-ipfs-api"
// )

// func main() {
// 	hash := "QmVQt8kV1HsNBBWUGYiKNup3j6EaCtBun2bMUcoG6iFCCd"
// 	sh := shell.NewLocalShell()
// 	RC, err := sh.Cat(hash)
// 	OnError(err)
// 	data, err := ioutil.ReadAll(RC)
// 	OnError(err)
// 	RC.Close()

// 	fmt.Printf("%s\n", data)
// }

// func OnError(err error) {
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		panic(err)
// 	}
// }
