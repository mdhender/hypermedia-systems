// Copyright (c) 2023 Michael D Henderson. All rights reserved.

package web1

//type Contacts struct {
//	web1 []Contact
//}
//
//func (c *Contacts) All() *Contacts {
//	if c == nil {
//		return &Contacts{}
//	}
//	return c
//}
//
//func (c *Contacts) Dump(w io.Writer) {
//	if c == nil {
//		_, _ = fmt.Fprintf(w, "[web1] dump: nil\n")
//		return
//	} else if len(c.web1) == 0 {
//		_, _ = fmt.Fprintf(w, "[web1] dump: []\n")
//		return
//	}
//	for i, contact := range c.web1 {
//		_, _ = fmt.Fprintf(w, "[web1] dump: %d %+v\n", i+1, contact)
//	}
//}
//
//func (c *Contacts) Search(search string) *Contacts {
//	s := &Contacts{}
//	if c == nil {
//		return s
//	}
//	matches := func(a, b string) bool {
//		return strings.EqualFold(a, b)
//	}
//	for _, contact := range c.web1 {
//		if matches(search, contact.First) {
//			s.web1 = append(s.web1, contact)
//		} else if matches(search, contact.Last) {
//			s.web1 = append(s.web1, contact)
//		} else if matches(search, contact.Phone) {
//			s.web1 = append(s.web1, contact)
//		} else if matches(search, contact.Email) {
//			s.web1 = append(s.web1, contact)
//		}
//	}
//	return s
//}
//
//type Contact struct {
//	Id     int
//	First  string
//	Last   string
//	Phone  string
//	Email  string
//	Errors struct {
//		Record error
//		First  error
//		Last   error
//		Phone  error
//		Email  error
//	}
//}
//
//func NewContacts(web1 ...Contact) *Contacts {
//	c := &Contacts{}
//	for _, contact := range web1 {
//		c.web1 = append(c.web1, contact)
//	}
//	return c
//}
//
//func NewContact(id int, first, last, phone, email string) Contact {
//	return Contact{
//		Id:    id,
//		First: first,
//		Last:  last,
//		Phone: phone,
//		Email: email,
//	}
//}
//
//func (c Contact) OK() bool {
//	return c.Errors.Record == nil && c.Errors.First == nil && c.Errors.Last == nil && c.Errors.Phone == nil && c.Errors.Email == nil
//}
//
//func (c Contact) Save() Contact {
//	if !c.OK() {
//		c.Errors.Record = fmt.Errorf("invalid contact")
//		return c
//	}
//	return c
//}
