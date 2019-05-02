// 10 国际化和本地化
// 为了适应经济的全球一体化，作为开发者，需要开发出支持多国语言、国际化的web应用，即同样的页面在不同的语言环境下需要显示不同的效果，也就是说应用程序在运行时能够根据请求所来自的地域与语言的不同而显示不同的用户界面。这样，当需要在应用程序中添加对新的语言的支持时，无需修改应用程序的代码，只需要增加语言包即可实现
// 国际化与本地化(Internationalization and localization，通常用il8n和L10N表示)，国际化是将针对某个地区设计的程序进行重构，以使它能够在更多时区使用，本地化是指在一个面向国际化的程序中增加对新地区的支持
// 目前，go的标准包没有提供对il8n的支持，但有一些比较简单的第三方实现，这一章将实现一个go-il8n库，用来支持go的il8n
// 所谓的国际化：就是根据特定的locale信息，提取与之相应的字符串或其他一些东西(比如时间和货币的格式)等等。这涉及到三个问题：
//	1. 如何确定locale
//	2. 如何保存与locale相关的字符串或其他信息
//	3. 如何根据locale提取字符串和其他相应的信息
// 在第一小节里，将介绍如何设置正确的locale以便让访问站点的用户能够获得与其语言相应的页面。
// 第二小节将介绍如何处理或存储字符串、货币、时间日期等与locale相关的信息，
// 第三小节将介绍如何实现国际化站点，即如何根据不同locale返回不同合适的内容