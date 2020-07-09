# go与机器学习

# 第一部分 机器学习里的数据分析

本章我们将了解一下go的数据分析，并重点讲解其在机器学习流程中的应用。

- 第一章, 收集数据
- 第二章, 矩阵，概率和统计
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

# CSV
> CSV是一种常见的数据保存形式

从csv文件中获取数据

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

# json

通常，当易用性是数据交换的主要目标时，将使用JSON。由于JSON是人类可读的，因此如果发生故障，很容易进行调试。请记住，在使用Go处理数据时，我们希望保持数据处理的完整性，并且该过程的一部分是确保在可能的情况下我们的数据可解释和可读

> 在当今世界中，大多数数据都是通过Web进行访问的，并且大多数工程组织都实现了一定数量的微服务，我们将相当频繁地遇到JSON格式的数据。我们可能仅在从API中提取一些随机数据时需要处理它，或者它实际上可能是驱动我们的分析和机器学习工作流程的主要数据格式。


# Parsing JSON

使用标准库中的encoding/json解析json

# SQL-like databases
> 当然还有数据库中的数据 这里使用postgres

- 例子1：检查数据库是否连接
- 例子2：查询
- 例子3：修改

# Caching
> 有时，我们的机器学习算法会受到来自外部来源（例如API）的数据的训练和/或指定的预测输入。也就是说，不是运行我们的建模或分析的应用程序本地的数据。此外，我们可能拥有经常访问的各种数据集，可能很快就会再次访问，或者可能需要在应用程序运行时使其可用。

在这些情况中的至少某些情况下，将数据缓存在内存中或在应用程序正在运行的地方本地嵌入数据可能是有意义的。例如，如果您经常访问政府API（通常具有高延迟）以获取人口普查数据，则可以考虑维护正在使用的人口普查数据的本地或内存中的缓存，以便避免不断地与人口普查数据联系。

# 内存缓冲
>github.com/patrickmn/go-cache

使用此第三方库，我们可以创建键和相应值的内存中缓存。我们甚至可以在缓存中为特定的键值对指定特性属性，例如生存时间。

# 数据库缓冲 将数据缓冲到硬盘
> 当然redis是国内的公司更多的选择，但这里将使用国际网络上流行的文件性数据库

我们刚刚看到的缓存在内存中。也就是说，缓存的数据存在，并且在您的应用程序运行时可以访问，但是一旦应用程序退出，数据就会消失。在某些情况下，您可能希望在重新启动或退出应用程序时保留缓存的数据。您可能还希望备份缓存，这样就不必在没有相关数据缓存的情况下从头启动应用程序。

# Data versioning

如前所述，机器学习模型会根据您使用的训练数据，参数的选择和输入数据产生截然不同的结果。出于协作，创意和合规性的原因，能够复制结果至关重要：

- Collaboration: 
> 对于一个团队，很难找到同时精通数据科学和精通机器学习的人（也就是说，在数据科学和机器学习的每个领域中都有知识和能力的人）。我们需要让我们的同事审查和改进我们的工作，如果他们无法重现我们的模型结果和分析，这是不可能获得正确的模型的。

- Creativity: 
> 我们不能相信自己会永远记住我们的推理和逻辑，尤其是在处理机器学习工作流程时。我们需要准确跟踪我们正在使用的数据，创建的结果以及创建方式。这是我们将能够不断改进模型和技术的唯一方法。

- Compliance
> 最后，我们可能很快就无法选择机器学习中的数据版本控制和可再现性。世界各地都在通过法律（例如，欧盟的通用数据保护条例（GDPR）），该法律赋予用户使用算法做出决定的解释权。如果我们没有一种可靠的方式来跟踪我们正在处理的数据和产生的结果，我们根本就不希望遵守这些裁决。

# 第二章 矩阵，概率和统计
>对概率和统计量的基本了解将使我们能够将某些算法与相关问题进行匹配，了解我们的数据和结果，并对我们的数据进行必要的转换。然后，矩阵和一些线性代数将使我们能够正确表示我们的数据，并实现优化，最小化和基于矩阵的转换。

如果您对数学或统计有些生疏，请不要担心太多。我们将在这里介绍一些基础知识，并向您展示如何以编程方式使用相关的统计量度和矩阵技术，这些将在本书的后面部分中加以利用。这不是一本有关统计，概率和线性代数的书。但是，要真正精通机器学习，您应该花时间更深入地学习这些主题。

在本章中，我们将介绍以下主题：
- 矩阵和向量
- 分布和统计指标
- 概率测度
- 假设检验

# 矩阵和向量
>实际上，许多机器学习算法都可以归结为一系列针对矩阵的迭代操作。

我们使用的第三方库为 github.com/gonum

# 向量
>向量是数字的有序集合，以行（从左到右）或列（上和下）排列。向量中的每个数字称为分量。例如，这可能是代表我们公司销售的数字的集合，或者可能是代表温度的数字的集合。

当然，我们使用切片来表示这些有序的数据集合，如下所示：

    // Initialize a "vector" via a slice.
    myVector := make([]float64, 0) 
    // Add a couple of components to the vector.
    myVector = append(myVector, 1.0)
    myVector = append(myVector, 2.2)
    // Output the results to stdout.
    fmt.Println(myVector)
    
但是，切片并不能真正代表行或列的概念，我们仍然需要在切片之上进行各种矢量运算。

    gonum.org/v1/gonum/floats
    gonum.org/v1/gonum/mat
 
如前所述，使用向量必须使用某些向量或特定于矩阵的运算和规则。例如，如何将向量相乘？我们如何知道两个向量是否相似？ gonum.org/v1/gonum/floats和gonum.org/v1/gonum/mat都提供矢量/切片操作的内置方法和功能，例如点积，排序和距离。我们将在这里不介绍所有功能，因为有很多内容，但是我们可以对使用向量的方式有一个大致的了解。
 
- 例子1：使用slice
- 例子2：使用gonum
- 例子3：使用floats进行矢量运算
- 例子4：使用mat进行矢量运算
