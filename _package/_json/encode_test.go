package json_test

import (
	"encoding/json"
	"fmt"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

// 临时忽略字段
func TestIgnoreField(t *testing.T) {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	b, _ := json.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User: &User{},
	})
	fmt.Println(string(b))

	b2, _ := jsoniter.Marshal(struct {
		*User
		Password bool `json:"password,omitempty"`
	}{
		User: &User{},
	})
	fmt.Println(string(b2))
}

// 临时添加字段
func TestAddField(t *testing.T) {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	b, _ := json.Marshal(struct {
		*User
		Token string `json:"token"`
	}{
		User:  &User{},
		Token: "test",
	})
	fmt.Println(string(b))

	b2, _ := jsoniter.Marshal(struct {
		*User
		Token string `json:"token"`
	}{
		User:  &User{},
		Token: "test",
	})
	fmt.Println(string(b2))
}

// 粘合两个struct
func TestMergeStruct(t *testing.T) {

	type BlogPost struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	}

	type Analytics struct {
		Visitors  int `json:"visitors"`
		PageViews int `json:"page_views"`
	}

	b, _ := json.Marshal(struct {
		*BlogPost
		*Analytics
	}{&BlogPost{}, &Analytics{}})
	fmt.Println(string(b))

	b2, _ := jsoniter.Marshal(struct {
		*BlogPost
		*Analytics
	}{&BlogPost{}, &Analytics{}})
	fmt.Println(string(b2))
}

// 一个json切分成两个struct
func TestSpliteTo2Struct(t *testing.T) {
	type BlogPost struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	}

	type Analytics struct {
		Visitors  int `json:"visitors"`
		PageViews int `json:"page_views"`
	}

	var post BlogPost
	var analytics Analytics

	s := `{
  "url": "attila@attilaolah.eu",
  "title": "Attila's Blog",
  "visitors": 6,
  "page_views": 14
}`
	_ = json.Unmarshal([]byte(s), &struct {
		*BlogPost
		*Analytics
	}{&post, &analytics})

	fmt.Printf("type: %T, value: %+v\n", post, post)
	fmt.Printf("type: %T, value: %+v\n", analytics, analytics)

	var post2 BlogPost
	var analytics2 Analytics
	_ = jsoniter.Unmarshal([]byte(s), &struct {
		*BlogPost
		*Analytics
	}{&post2, &analytics2})

	fmt.Printf("type: %T, value: %+v\n", post2, post2)
	fmt.Printf("type: %T, value: %+v\n", analytics2, analytics2)
}

// 修改字段名字
func TestChangeFieldName(t *testing.T) {
	type Value int
	type CacheItem struct {
		Key    string `json:"key"`
		MaxAge int    `json:"cacheAge"`
		Value  *Value `json:"cacheValue"`
	}

	v := Value(10086)

	item := CacheItem{
		Key:    "test",
		MaxAge: int(v),
		Value:  &v,
	}
	b, _ := json.Marshal(struct {
		*CacheItem

		// 添加字段使用老的field name
		OmitMaxAge int `json:"cacheAge,omitempty"`
		OmitValue  int `json:"cacheValue,omitempty"`

		// 字段重命名
		MaxAge int    `json:"max_age"`
		Value  *Value `json:"value"`
	}{
		CacheItem:  &item,
		OmitMaxAge: 1,
		OmitValue:  2,
		MaxAge:     item.MaxAge,
		Value:      &v,
	})
	fmt.Printf(string(b))
}

// 添加扩展对象
func TestExtend(t *testing.T) {
	type Value int
	type CacheItem struct {
		Key    string `json:"key"`
		MaxAge int    `json:"cacheAge"`
		Value  *Value `json:"cacheValue"`
	}

	type Extend struct {
		BoyFriendsName string `json:"b_f_name"`
	}

	v := Value(10086)

	item := CacheItem{
		Key:    "test",
		MaxAge: int(v),
		Value:  &v,
	}
	extend := Extend{
		"哈哈",
	}
	b, _ := json.Marshal(struct {
		*CacheItem
		*Extend
	}{
		CacheItem: &item,
		Extend:    &extend,
	})
	fmt.Printf(string(b))
}
