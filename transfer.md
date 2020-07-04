Section 1: Analysis in Machine Learning Workflows

In this section, we will get a solid understanding of how to parse and organize data within a Go program, with an emphasis on handling that data in a machine learning workflow.

This section will contain the following chapters:

- Chapter 1, Gathering and Organizing Data
- Chapter 2, Matrices, Probability, and Statistics
- Chapter 3, Evaluating and Validating

# 1 Gathering and Organizing Data

Machine learning in general involves a series of steps, out of which the process of gathering and cleaning data consumes a lot of time. Polls have shown that 90% or more of a data scientist's time is spent gathering data, organizing it, and cleaning it—not training or tuning their sophisticated machine learning models. Why is this? Isn't the machine learning part the fun part? Why do we need to care so much about the state of our data?

Not all types of data are appropriate when using certain types of models. For example, certain models do not perform well when we have high-dimensional data (for example, text data), and other models assume that variables are normally distributed, which is definitely not always the case. Thus, we must take care to gather data that fits our use case and make sure that we understand how our data and models will interact.

Another reason why gathering and organizing data consumes so much of a data scientist's time is that data is often messy and hard to aggregate. In most organizations, data might be housed in various systems and formats, and have various access control policies. To form a training or test set or to supply variables to a model for predictions, we will likely need to deal with various formats of data, such as CSV, JSON, and database tables, and we will likely need to transform individual values. Common transformations include missing value analysis, parsing date times, converting categorical data to numerical data, normalizing values, and applying some function across values. Scraping the data from the web has emerged as an important data source and a lot of data-driven institutions are relying on it for adding it to their repositories.

Even though much of this book will be focused on various modeling techniques, you should always consider data gathering, parsing, and organization as a – or maybe the – key component of a successful data science project. If this part of your project is not carefully developed with a high level of integrity, you are setting yourself up for trouble in the long run.


From this chapter, readers will be able to learn the different data handling techniques using
Golang with guided code covering the following topics:

- Handling varied data forms—CSV, JSON, and SQL databases
- Web scraping
- Caching
- Data versioning

CSV files might not be a go-to format for big data, but as a data scientist or developer working in machine learning, you are sure to encounter this format. You might need a mapping of zip codes to latitude/longitude and find this as a CSV file on the internet, or you may be given sales figures from your sales team in a CSV format. In any event, we need to understand how to parse these files. The main package that we will utilize in parsing CSV files is encoding/csv from Go's standard library. However, we will also discuss a couple of packages that allow us to quickly manipulate or transform CSV data, github.com/go-gota/gota/dataframe and go-hep.org/x/hep/csvutil.
