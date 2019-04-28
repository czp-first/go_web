// 7.1 XML处理
// 本节是介绍如何用go来编程编码XML文件相关的知识
// 假如你是一名运维人员，你为你所管理的所有服务器生成了如下内容的xml的配置文件：
<?xml version="1.0" encoding="utf-8"?> 
<servers version="1">
	<server> 
		<serverName>Shanghai_VPN</serverName> 
		<serverIP>127.0.0.1</serverIP>
	</server> 
	<server>
		<serverName>Beijing_VPN</serverName>
		<serverIP>127.0.0.2</serverIP>
	 </server>
</servers>
// 上面的XML文档描述了两个服务器的信息，包含了服务器名和服务器的IP信息，接下来的go例子以此XML描述的信息进行操作

// 解析XML
// 通过xml包的Unmarshal函数来达到目的
// func Unmarshal(data []byte, v interface{}) error
// data接收的是XML数据流，v是需要输出的结构，定义为interface，也就是可以把XML转换为任意的格式。这里主要介绍struct的转换，因为struct和XML都有类似树结构的特征

// parse_xml.go中，将xml文件解析成对应的struct对象是通过xml.Unmarshal来完成的，这个过程是如何实现的？可以看到struct定义后面对了类似于xml:"serverName"这样的内容，这个是struct的一个特性，它们被称为struct tag，它们是用来辅助反射的。
// 看一下Unmarshal的定义:
// func Unmarshal(data []byte, v interface{}) error
// 函数定义了两个参数，第一个是XML数据流，第二个是存储的对应类型，目前支持struct、slice和string，XML包内部采用了反射来进行数据的映射，所以v里面的字段必须是导出的。Unmarshal解析的时候XML元素和字段名对应起来的呢？这是有一个优先级读取流程的，首先会读取struct tag，如果没有，那么就会对应字段名。必须注意一点的是解析的时候tag、字段名、XML元素都是大小写敏感的，所以必须一一对应
// go的反射机制，可以利用这些tag信息来将来自XML文件中的数据反射成对应的struct对象，关于反射如何利用struct tag的更多内容请参阅reflect中的相关内容
// 解析XML到struct的时候遵循如下的规则：
//	如果struct的一个字段是string或者[]byte类型且他的tag含有",innerxml",Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上，如例子中Description定义，最后的输出是
//	ShanghaiVPN 127.0.0.1 BeijingVPN 127.0.0.2
//	如果struct中有一个叫做XMLName，且类型为xml.Name字段，那么在解析的时候就会保存这个element的名字到该字段，如上面例子中的servers
//	如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析的时候就会把相应的element值赋值给该字段，如上servername和serverip定义
//	如果某个struct字段的tag定义中含有",attr"，那么解析的手就会将该结构所对应的element的与字段同名的属性的值赋值给该字段，如上version定义
//	如果某个struct字段的tag定义 型如"a>b>c"，则解析的时候，会将xml结构a下面的b下面的c元素赋值给该字段
//	如果某个struct字段的tag定义了"-"，那么不会为该字段解析匹配任何xml数据
//	如果struct字段后面的tag定义了",any"，如果他的子元素在不满足其他的规则的时候就会匹配到这个字段
//	如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，这个字段的类型可能是[]byte或string，如果没有这样的字段存在，那么注释将会被抛弃
//	注意：为了正确解析，go的xml包要求struct定义中的所有字段必须是可导出的(即首字母大写)

// 输出XML
// 假如不是要解析如上所示的XML文件，而是生成它，那么在go中该如何实现
// xml包中提供了Marshal和MarshalIndent两个函数，来满足需求。这两个函数的主要区别是第二个函数会增加前缀和缩进，函数的定义如下所示：
// func Marshal(v interface{}) ([]byte, error)
// func MarshalIndent(v interface{}, prefix, indent string) ([] byte, error)
// 两个函数第一个参数是用来生成XML的结构定义类型数据，都是返回生成的XML数据流 

// 例子：output_xml.go

// 和之前定义的文件的格式一摸一样，之所以会有os.Stdout.Write([]byte(xml.Header))这句代码出现，是因为xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头的，为了生成正确的xml文件，使用了xml包预定义的Header变量
// Marshal函数接收的参数v是interface{}类型的，即它可以接受任意类型的参数，那么xml包，根据什么规则生成相应的XML文件呢？
//	如果v是array或者slice，那么输出每一个元素，类似value
//	如果v是指针，那么会Marshal指针指向的内容，如果指针为空，什么都不输出
//	如果v是interface，那么就处理interface所包含的数据
//	如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息
// 生成的XML文件中的element的名字又是根据什么决定的呢？元素名按照如下优先级从struct中获取：
//	如果v是struct，XMLName的tag中定义的名称
//	类型为xml.Name的名叫XMLName的字段的值
//	通过struct中字段的tag来获取
//	通过struct的字段名用来获取
//	marshal的类型名称
// 应该如何设置struct中字段的tag信息以控制最终xml文件的生成
//	XMLName不会被输出
//	tag中含有"-"的字段不会输出
//	tag中含有"name,attr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述
//	tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似上一条，只是这个name默认是字段名了
//	tag含有",chardata"，输出为xml的character data而非element
//	tag含有",innerxml", 将会被原样输出，而不会进行常规的编码过程
//	tag含有",comment"，将被当作xml注释来输出，而不会进行常规的编码过程，字段值不能含有"--"字符串
//	tag含有"omitemoty",如果该字段的值为空值那么该字段就不会被输出到XML，空值包括:false、0、nil指针或nil接口，任何长度为0的array,slice，map和string
//	tag中含有"a>b>c", 那么就会循环输出三个元素a包含b，b包含c，例如如下代码就会输出
FirstName string `xml:"name>first"`
LastName string `xml:"name>last"`

<name>
	<first>Asta</first>
	<last>Xie </last>
</name>

// XML的所有操作都是通过struct tag来实现的，所以学会对struct tag的运用变得非常重要。












































