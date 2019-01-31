library(plotly)

Labels <- c("Average", "99% Line", "Throughput")

#128mb RAM
Go128Mb <- c(24, 70, 463.45)
Java128Mb <- c(24, 53, 468.21)
Python128Mb <- c(39, 78,346.75)

data128 <- data.frame(Labels, Go128Mb , Java128Mb , Python128Mb)

p128 <- plot_ly(data128, x = ~Labels, y = ~Go128Mb, type = 'bar', name = 'Go 128Mb') %>%
  add_trace(y = ~Java128Mb, name = 'Java 128Mb') %>%
  add_trace(y = ~Python128Mb, name = 'Python 128Mb') %>%
  layout(yaxis = list(title = 'Count'), barmode = 'group')

#32Mb RAM
Go32Mb <- c(23, 52, 477.53)
Python32Mb <- c(39, 86, 351.98)

data32 <- data.frame(Labels, Go32Mb, Python32Mb)

p32 <- plot_ly(data, x = ~Labels, y = ~Go32Mb, type = 'bar', name = 'Go 32Mb') %>%
  add_trace(y = ~Python32Mb, name = 'Python 32Mb') %>%
  layout(yaxis = list(title = 'Count'), barmode = 'group')





