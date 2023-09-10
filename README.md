### golang-restful-api-framework
1. model，在model中增加do，对应设计的table
2. dao，在dao中增加crud操作，即对1种的table进行操作
3. serivce，在service中增加对请求的操作，及对dao获得结果的操作
4. controller，在controller中承接请求，同时调用3中的service获得结果，进行响应
5. setup 进行注册
    1. init全局变量
    2. init_wire_inject处理依赖
6. router 进行注册路由
7. main，调用6 加载路由