# go与机器学习

# 第一部分 机器学习里的数据分析

本章我们将了解一下go的数据分析，并重点讲解其在机器学习流程中的应用。

- 第一章, 收集数据
- 第二章, Matrices, Probability, and Statistics
- 第三章, Evaluating and Validating

# 第一章 收集数据

在传统的机器学习中，所耗费时间最长的是数据的收集和清洗。在机器学习中，数据的选取有重要意义，但是数据的来源往往混乱而且复杂。为了形成训练或测试集或向模型提供变量以进行预测，我们可能需要处理各种格式的数据，例如CSV，JSON和数据库表，并且可能需要转换单个值。

常见的转换包括缺失值分析，解析日期时间，将分类数据转换为数值数据，对值进行归一化以及跨值应用某些函数。从网络中收集数据已成为一种重要的数据源，许多数据驱动的机构都依赖于将其添加到其存储库中。常见的转换包括缺失值分析，解析日期时间，将分类数据转换为数值数据，对值进行归一化以及跨值应用某些函数。从网络中收集数据已成为一种重要的数据源，许多数据驱动的机构都依赖于将其添加到其存储库中。


本章所涉及到的数据来源
- Handling varied data forms—CSV, JSON, and SQL databases
- Web scraping
- Caching
- Data versioning

#1.1 处理数据

- CSV
> CSV是一种常见的数据保存形式

# 从csv文件中获取数据

文件 https:/​/​archive.​ics.​uci.​edu/​ml/​datasets/​iris

该CSV文件包括四个花朵测量值的浮动列和一个带有相应花朵种类的字符串列

- 例子1：读取文件
- 例子2：处理文件中每行的个数是否完整
> Reader.FieldsPerRecord规定每行中含有的元素个数

- 例子3：处理文件中每行各个元素的数据类型是否复合预期
- 例子4：使用go-gota解析数据
- 例子5：使用go-gota处理数据

# Web scraping
> 网络爬取

在多种情况下收集数据可能很有用，例如在网站不提供API时，或者您需要以编程方式解析和提取Web内容时（例如，抓取Wikipedia表）。

- github.com/PuerkitoBio/goquery
> 类似jquery
- github.com/anaskhan96/soup
> 类似BeautifulSoup(python包)

