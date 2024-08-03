# coding:utf-8

import os
from flask import Flask, Blueprint, render_template, request
import urllib3, json, base64, time, hashlib
from datetime import datetime
from utils import sm2

urllib3.disable_warnings()

hostname = '127.0.0.1'
chain_api_port = [ "8001", "8002" ]

demo_app = Blueprint('demo', __name__)


# 接口演示
@demo_app.route("/chain_demo", methods=["GET"])
def demo_get():
    return render_template('demo.html', chain_api_port=chain_api_port)

@demo_app.route("/chain_demo", methods=["POST"])
def demo_post():
    #print(dict(request.form))

    body_data = process_demo(dict(request.form))
    #print("body_data", body_data)
    api_url, params, status, rdata, timespan = \
        call_api(hostname, request.form['port'], request.form['url'], body_data)
    return render_template('result.html', 
        result=rdata, status=status, 
        timespan=timespan, params=params, api_url=api_url)


# 生成参数字符串
def gen_param_str(param1):
    param = param1.copy()
    name_list = sorted(param.keys())
    if 'data' in name_list: # data 按 key 排序, 中文不进行性转义，与go保持一致
        param['data'] = json.dumps(param['data'], sort_keys=True, ensure_ascii=False, separators=(',', ':'))
    return '&'.join(['%s=%s'%(str(i), str(param[i])) for i in name_list if str(param[i])!=''])


# 调用接口
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

    body_str = json.dumps(body)
    print(body)

    pool = urllib3.PoolManager(num_pools=2, timeout=180, retries=False)

    host = 'http://%s:%s'%(hostname, port)

    if action=='test':
        url = host + f"/api/test"
    else:
        url = host + f"/api/r1/{action}"

    print(url)

    start_time = datetime.now()
    r = pool.urlopen('POST', url, body=body_str)

    print(r.status)
    if r.status==200:
        rdata = json.dumps(json.loads(r.data.decode('utf-8')), ensure_ascii=False, indent=4)
    else:
        rdata = r.data

    body2 = json.dumps(body, ensure_ascii=False, indent=4)
    return url, body2, r.status, rdata, \
        '{!s}s'.format(datetime.now() - start_time)



def process_demo(form_data):
    body_data = form_data.copy()
    body_data.pop('url')
    body_data.pop('port')

    # 查询
    if form_data['url'] == 'q/user/list':
        body_data['page'] = int(body_data['page'])
        body_data['limit'] = int(body_data['limit'])
        if len(form_data['status'])==0:
            body_data.pop('status')

    elif form_data['url'] in ('q/user/info', 'q/user/verify', 'q/bank/balance', 'q/kv/show', 'q/block/tx', 'q/block/height'):
        pass

    elif form_data['url'] == 'q/kv/list':
        body_data['page'] = int(body_data['page'])
        body_data['limit'] = int(body_data['limit'])
        if len(form_data['owner_addr'])==0:
            body_data.pop('owner_addr')

    elif form_data['url'] == 'q/block/txs':
        body_data['page'] = int(body_data['page'])
        body_data['limit'] = int(body_data['limit'])
        if len(form_data['creator_addr'])==0:
            body_data.pop('creator_addr')
        if len(form_data['tx_action'])==0:
            body_data.pop('tx_action')

    elif form_data['url'] in ('q/exchange/ask/list', 'q/exchange/reply/list'):
        body_data['page'] = int(body_data['page'])
        body_data['limit'] = int(body_data['limit'])
        if len(form_data['asker_addr'])==0:
            body_data.pop('asker_addr')
        if len(form_data['replier_addr'])==0:
            body_data.pop('replier_addr')
        if len(form_data['uuid'])==0:
            body_data.pop('uuid')

    elif form_data['url'] == 'q/exchange/reply/show':
        body_data['reply_id'] = int(body_data['reply_id'])
        body_data['decrypt'] = body_data['decrypt']=='true'

    elif form_data['url'] == 'q/post/recv/show':
        body_data['post_id'] = int(body_data['post_id'])
        body_data['decrypt'] = body_data['decrypt']=='true'

    elif form_data['url'] in ('q/post/sent/list', 'q/post/timeout/list', 'q/post/recv/list'):
        body_data['page'] = int(body_data['page'])
        body_data['limit'] = int(body_data['limit'])
        if len(form_data['sender_addr'])==0:
            body_data.pop('sender_addr')
        if len(form_data['target_addr'])==0:
            body_data.pop('target_addr')
        if len(form_data['uuid'])==0:
            body_data.pop('uuid')

    # 交易
    elif form_data['url'] in ('tx/user/new', 'tx/user/update'):
        if len(form_data['name'])==0:
            body_data.pop('name')
        if len(form_data['acc_no'])==0:
            body_data.pop('acc_no')
        if len(form_data['address'])==0:
            body_data.pop('address')
        if len(form_data['phone'])==0:
            body_data.pop('phone')
        if len(form_data['ref'])==0:
            body_data.pop('ref')

    elif form_data['url'] in ('tx/user/audit', 'tx/kv/delete', 'tx/exchange/ask', 'tx/post/send', 'tx/post/ask'):
        pass

    elif form_data['url'] in ('tx/kv/new', 'tx/kv/update'):
        body_data['crypto'] = body_data['crypto']=='true'

    elif form_data['url'] == 'tx/exchange/reply':
        body_data['ask_id'] = int(body_data['ask_id'])
        body_data['reply'] = body_data['reply']=='true'
        if len(form_data['payload'])==0:
            body_data.pop('payload')

    elif form_data['url'] == 'tx/post/reply':
        body_data['ask_post_id'] = int(body_data['ask_post_id'])
        body_data['reply'] = body_data['reply']=='true'
        if len(form_data['payload'])==0:
            body_data.pop('payload')

    return body_data
