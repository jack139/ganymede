# coding:utf-8
import sys
import urllib3, json, base64, time, hashlib
from datetime import datetime
from demo.utils import sm2
urllib3.disable_warnings()


# 生成参数字符串
def gen_param_str(param1):
    param = param1.copy()
    name_list = sorted(param.keys())
    if 'data' in name_list: # data 按 key 排序, 中文不进行性转义，与go保持一致
        param['data'] = json.dumps(param['data'], sort_keys=True, ensure_ascii=False, separators=(',', ':'))
    return '&'.join(['%s=%s'%(str(i), str(param[i])) for i in name_list if str(param[i])!=''])


def call_api(hostname, port, action, body_data):

    body = {
        'version'  : '1',
        'sign_type' : 'SM2', 
        'data'     : body_data,
    }

    secret = 'MjdjNGQxNGU3NjA1OWI0MGVmODIyN2FkOTEwYTViNDQzYTNjNTIyNSAgLQo='
    appid = '4fcf3871f4a023712bec9ed44ee4b709'
    unixtime = int(time.time())
    body['timestamp'] = unixtime
    body['appid'] = appid

    param_str = gen_param_str(body)
    sign_str = '%s&key=%s' % (param_str, secret)

    if body['sign_type'] == 'SHA256':
        sha256 = hashlib.sha256(sign_str.encode('utf-8')).hexdigest().encode('utf-8')
        signature_str =  base64.b64encode(sha256).decode('utf-8')
    else: # SM2
        signature_str = sm2.SM2withSM3_sign_base64(sign_str)

    #print(sign_str.encode('utf-8'))
    #print(sha256)
    #print(signature_str)

    body['sign_data'] = signature_str

    body = json.dumps(body)
    print(body)

    pool = urllib3.PoolManager(num_pools=2, timeout=180, retries=False)

    host = 'http://%s:%s'%(hostname, port)

    if action=='test':
        url = host + f"/api/test"
    else:
        url = host + f"/api/r1/{action}"

    print(url)

    start_time = datetime.now()
    r = pool.urlopen('POST', url, body=body)
    print('[Time taken: {!s}]'.format(datetime.now() - start_time))

    print(r.status)
    if r.status==200:
        #print(json.loads(r.data.decode('utf-8')))
        print(r.data.decode('utf-8'))
    else:
        print(r.data)



# 调用
def user_register():
    body_data = {
        'key_name' : 'test2',
        'user_type' : 'USR',
    }
    call_api(hostname, port, 'tx/user/new', body_data)


def user_audit():
    body_data = {
        'chain_addr' : 'saturn10fhk7pqehrjcz8maje4hvg2e200vwsdl05z9s0',
        'status' : 'ACTIVE',
    }
    call_api(hostname, port, 'tx/user/audit', body_data)


def kv_new():
    body_data = {
        'owner_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs',
        'key'   : 'k2',
        'value' : 'v1',
        'crypto' : 'plain',
    }
    call_api(hostname, port, 'tx/kv/new', body_data)


def kv_update(o, k, v):
    body_data = {
        'owner_addr' : o,
        'key'   : k,
        'value' : v,
        'crypto' : 'plain',
    }
    call_api(hostname, port, 'tx/kv/update', body_data)


def kv_show(o, k):
    body_data = {
        'owner_addr' : o,
        'key'   : k,
        'crypto' : 'plain',
    }
    call_api(hostname, port, 'q/kv/show', body_data)


def post_ask():
    body_data = {
        'asker_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs', # test2
        'replier_addr' : 'saturn1xx0a7m669fdujvcwux0zcrgw3fzy5a79a8e8cl', # test4
        'payload' : 'post ask payload 22222',
        'post_channel' : 'channel-0',
    }
    call_api(hostname, port, 'tx/post/ask', body_data)

def post_reply():
    body_data = {
        'asker_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs', # test2
        'replier_addr' : 'saturn1xx0a7m669fdujvcwux0zcrgw3fzy5a79a8e8cl', # test4
        'payload' : 'post reply content 5555',
        'post_channel' : 'channel-0',
        'ask_post_id' : 1,
        'reply' : True,
    }
    call_api(hostname, port, 'tx/post/reply', body_data)


def post_send():
    body_data = {
        'sender_addr' : 'saturn1dp6eqr44lcf4q6q8aaegkh73r3uc5d2068peal', # test5
        'target_addr' : 'saturn1xx0a7m669fdujvcwux0zcrgw3fzy5a79a8e8cl', # test4
        'payload' : 'post sent some 测试测试',
        'post_channel' : 'channel-0',
    }
    call_api(hostname, port, 'tx/post/send', body_data)


def exchange_ask():
    body_data = {
        'asker_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs', # test2
        'replier_addr' : 'saturn10fhk7pqehrjcz8maje4hvg2e200vwsdl05z9s0', # test3
        'payload' : 'ask payload 测试 2222',
    }
    call_api(hostname, port, 'tx/exchange/ask', body_data)

def exchange_reply():
    body_data = {
        'asker_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs', # test2
        'replier_addr' : 'saturn10fhk7pqehrjcz8maje4hvg2e200vwsdl05z9s0', # test3
        'payload' : 'reply test 测试',
        'ask_id' : 0,
        'reply' : False,
    }
    call_api(hostname, port, 'tx/exchange/reply', body_data)



if __name__ == '__main__':
    if len(sys.argv)<2:
        print("usage: python3 %s <host> <port>" % sys.argv[0])
        sys.exit(2)

    hostname = sys.argv[1]
    port = sys.argv[2]



    body_data = {

        #'chain_addr'   : 'saturn1dp6eqr44lcf4q6q8aaegkh73r3uc5d2068peal', # test5
        #'mystery' : "turkey such parrot never divorce dust cube twist entry climb weapon rotate million network cruise senior vintage ramp pull volcano prize inherit depth phrase", # test2
        #'mystery' : "empty question food laundry",
        #'positions' : "1 3 5 24",

        #'target_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs', # test2
        #'target_addr' : 'saturn10fhk7pqehrjcz8maje4hvg2e200vwsdl05z9s0', # test3
        #'target_addr' : 'saturn1xx0a7m669fdujvcwux0zcrgw3fzy5a79a8e8cl', # test4
        #'target_addr' : 'saturn1dp6eqr44lcf4q6q8aaegkh73r3uc5d2068peal', # test5
        #'post_id' : 1,
        #'decrypt' : True,

        #'asker_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs',
        #'replier_addr' : 'saturn10fhk7pqehrjcz8maje4hvg2e200vwsdl05z9s0',
        #"uuid" : "b687bb0b-8889-4c6d-a22e-a5d8068e125b",
        #'ask_id' : 0,
        #'reply_id' : 0,
        #'decrypt' : True,

        #'height' : '985',

        #'page' : 1,
        #'limit' : 10,

        #'txhash' : 'A83D56175119567F48EB00C005E7C9504D72731A97132BA8C8B9277DFA10001E',
        #'creator_addr' : 'saturn1s5a7468v2akv523m39wca2ptp8f9wh2xq2fzhs',
        #'tx_action' : 'kv/new',
    }

    #call_api(hostname, port, 'q/block/txs', body_data)

    #call_api(hostname, port, 'q/user/list', body_data)

    #call_api(hostname, port, 'q/post/sent/list', body_data)
    #call_api(hostname, port, 'q/post/recv/list', body_data)
    #call_api(hostname, port, 'q/post/timeout/list', body_data)
    #call_api(hostname, port, 'q/post/recv/show', body_data)

    #call_api(hostname, port, 'q/exchange/ask/list', body_data)
    #call_api(hostname, port, 'q/exchange/reply/list', body_data)
    #call_api(hostname, port, 'q/exchange/reply/show', body_data)



    #user_register()
    #user_audit()
    #kv_new()
    kv_update('saturn1p4dgmrayu8lt7zspvdjqjhjqz4rgcywytjmrr6', 'k1', 'v1_new11111')
    #time.sleep(5)
    kv_update('saturn1p4dgmrayu8lt7zspvdjqjhjqz4rgcywytjmrr6', 'k2', 'v2_new22222')
    #kv_update('saturn140fwmp5drgrfk0n7503vje4xd4zqf6qn4akltm', 'k3', 'v3_new')

    kv_show('saturn1p4dgmrayu8lt7zspvdjqjhjqz4rgcywytjmrr6', 'k1')
    kv_show('saturn1p4dgmrayu8lt7zspvdjqjhjqz4rgcywytjmrr6', 'k2')

    #post_send()
    #post_ask()
    #post_reply()

    #exchange_ask()
    #exchange_reply()

    

    '''
        '/api/test'
        '/api/r1/tx/user/new'
        '/api/r1/tx/user/update'
        '/api/r1/tx/user/audit'
        '/api/r1/tx/kv/new'
        '/api/r1/tx/kv/update'
        '/api/r1/tx/kv/delete'
        '/api/r1/tx/exchange/ask'
        '/api/r1/tx/exchange/reply'
        '/api/r1/tx/post/send'
        '/api/r1/tx/post/ask'
        '/api/r1/tx/post/reply'

        '/api/r1/q/block/height'
        '/api/r1/q/block/tx'
        '/api/r1/q/bank/balance'
        '/api/r1/q/user/info'
        '/api/r1/q/user/list'
        '/api/r1/q/user/verify'
        '/api/r1/q/kv/show'
        '/api/r1/q/kv/list'
        '/api/r1/q/exchange/ask/list'
        '/api/r1/q/exchange/ask/show'
        '/api/r1/q/exchange/reply/list'
        '/api/r1/q/exchange/reply/show'
        '/api/r1/q/post/sent/list'
        '/api/r1/q/post/timeout/list'
        '/api/r1/q/post/recv/list'
        '/api/r1/q/post/recv/show'
    '''
