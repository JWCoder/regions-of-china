# Regions of China (中国省市地区和行政区划数据)
The administrative divisions of China, provide CSV, SQL, text, Excel and XML data format.

Latest upated time: 2016-07-31

# Data format
* regions-of-china.csv  CSV format.
* regions-of-china.sql  SQL format.
* regions-of-china.txt  Text format.
* regions-of-china.xlsx  Excel format.
* regions-of-china.xml  XML format.

# Table structure
```sql
CREATE TABLE `region` (
  `district_id` int(10) unsigned NOT NULL,
  `district_name` varchar(128) NOT NULL,
  `city_id` int(10) unsigned NOT NULL,
  `city_name` varchar(128) NOT NULL,
  `province_id` int(10) unsigned NOT NULL,
  `province_name` varchar(128) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

# 说明
中国行政区划数据. 包括省(直辖市)、市(县)和区, 数据来自"中华人民共和国统计局"官方网站. 数据截止2016年07月31日.

# References
* https://en.wikipedia.org/wiki/List_of_regions_of_the_People%27s_Republic_of_China
* https://en.wikipedia.org/wiki/Administrative_divisions_of_China
* https://en.wikipedia.org/wiki/Provinces_of_China
* http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201703/t20170310_1471429.html
