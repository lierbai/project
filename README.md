|--routers     路由层

|--controllers 控制器层,负责逻辑判断和返回结果

|--service     服务层,主要是处理控制层传入的数据并进行业务处理

|--Dao         数据访问层 是服务层获取数据的接口包

|--repository  数据仓库层,把数据库和redis和其他存储都放在这个包下

|--component   组件包,是主要四层的补充,里面一般放不确定那层需要调用的东西

|--conf        配置包

|--conn        连接包,包括redis、mysql、id生成器的初始化链接

|--constname   常量包,包括默认值,业务响应码,业务响应信息,redis的Key

|--filter      过滤器,负责非法请求的拦截、非法词汇拦截、接口基本鉴权、身份鉴权等

|--initialize  初始化包,当使用多存储的时候，需要先把数据库的基本数据初始化内存数据库中，比如redis

|--middleware  中间件包，比如jwt等

|--util        工具包,主要是一些小工具的包