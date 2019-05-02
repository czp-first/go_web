// 10.3 国际化站点
// 前面小节介绍了如何处理本地化资源，即Locale一个相应的配置文件，那么如果处理多个本地化资源呢？而对于一些我们经常用到的例如：简单的文本翻译、时间日期、数字等如何处理呢？

// 管理多个本地包
// 在开发一个应用的时候，首先要决定是只支持一种语言，还是多种语言，如果要支持多种语言，则需要制定一个组织结构，以方便将来更多语言的添加。为此设计如下:Locale有关的文件放置在config/locales下，假设要支持中文和英文，那么需要在这个文件夹下放置en.json和zh.json。大概的内容如下所示
// zh.json
// en.json
// 为了支持国际化，在此使用了一个国际化相关的包--go-il8n(https://github.com/astaxie/go-il8n),首先想go-il8n注册config/locales这个目录，以加载所有的locale文件
Tr := il8n.NewLocale()
Tr := LoadPath("config/locales")
// 这个包使用起来很简单，你可以通过下面的方式进行测试：
fmt.Println(Tr.Translate("submit"))
// 输出Submit
Tr.SetLocale("zn")
fmt.Println(Tr.Translate("submit"))

// 自动加载本地包
// 上面介绍了如何自动加载自定义语言包，其实go-il8n库已经预加载了很多默认的格式信息，例如时间格式、货币格式，用户可以在自定义配置时改写这些默认配置，请看下面的处理过程
// 加载默认配置文件，这些文件都放在go-il8n/locales下面
// 文件命名zh.json、en-json、en-US.json等，可以不断的扩展支持更多的语言
func (il *IL) loadDefaultTranslations(dirPath string) error {
	dir, err := os.Open(dirPath)
	if err := nil {
		return err
	}
	defer dir.Close()

	names, err := dir.Readdrinames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		fullPath := path.Join(dirPath, name)

		fi, err := os.Stat(fullPath)
		if err != nil {
			return err
		}
		if fi.IsDir() {
			if err := il.loadTranslations(fullPath); err != nil {
				return err
			}
		} else if locale := il.matchingLocaleFromFileName(name); locale != "" {
			file, err := os.Open(fullPath)
			if err != nil {
				return err
			}
			defer file.Close()
			if err := il.loadTranslation(file, locale); err != nil {
				return err
			}
		}
	}
	return nil
}
// 通过上面的方法加载配置信息到默认的文件，这样就可以在没有自定义时间信息的时候执行如下的代码获取对应的信息:
// locale=zh情况下，执行如下代码
fmt.Println(Tr.Time(time.Now())
// 输出：2009年1月08日 星期四 20:37:58 CST

fmt.Println(Tr.Time(time.Now(), "long"))
// 输出：2009年1月08日

fmt.Println(Tr.Money(11.11))
// 输出：¥11.11

// template mapfunc
// 上面实现了多个语言包的管理和加载，而一些函数的实现是基于逻辑层的，例如："Tr.Translate"、"Tr.Time"、"Tr.Money"等，虽然在逻辑层可以利用这些函数把需要的参数进行转换后在模版层渲染的时候直接输出，但是如果想要在模版层直接使用这些函数该怎么实现呢？
// 不知你是否还记得，在前面介绍模版的时候说过:go的模版支持自定义模版函数，下面是实现的方便操作的mapfunc

// 	1. 文本信息
// 文本信息调用Tr.Translate来实现相应的信息转换，mapFunc的实现如下
func Il8nT(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	return Tr.Translate(s)
}
// 注册函数如下
t.Func(template.FuncMap{"T": Il8nT})
// 模版中使用如下:
{{.V.Submit | T}}

//	1. 时间日期
// 时间日期调用Tr.Time函数来实现相应的时间转换，mapFunc的实现如下
func Il8bTimeDate(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	return Tr.Time(s)
}
// 注册函数如下：
t.Funcs(template.FuncMap{"TD": Il8nTimeDate})
// 模版中使用如下
{{.V.Now | TD}}

// 1. 货币信息
// 货币调用Tr.Money函数来实现相应的时间转换，mapFunc的实现如下
func Il8nMoney(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if != ok {
		s = fmt.Sprint(args...)
	}
	return Tr.Money(s)
}
// 注册函数如下：
t.Func(template.FuncMap{"M": Il8nMoney})
// 模版中使用如下
{{.V.Money | M}}











