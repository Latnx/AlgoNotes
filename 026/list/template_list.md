package main
// type Element
func (e *Element) Next() *Element
func (e *Element) Prev() *Element

// type List
func New() *List
func (l *List) Back() *Element                             //最后一个元素
func (l *List) Front() *Element                            //第一个元素
func (l *List) Init() *List                                //初始化
func (l *List) InsertAfter(v any, mark *Element) *Element  //在某元素后插入
func (l *List) InsertBefore(v any, mark *Element) *Element //在某元素前插入
func (l *List) Len() int                                   //长度
func (l *List) MoveAfter(e, mark *Element)                 //移动某元素到mark后
func (l *List) MoveBefore(e, mark *Element)                //移动某元素到mark前
func (l *List) MoveToBack(e *Element)                      //移动到最后
func (l *List) MoveToFront(e *Element)                     //移动到开始
func (l *List) PushBack(v any) *Element                    //添加到最后
func (l *List) PushBackList(other *List)                   //添加list到最后
func (l *List) PushFront(v any) *Element                   //添加到开始
func (l *List) PushFrontList(other *List)                  //添加list到开始
func (l *List) Remove(e *Element) any                      //移除某个元素