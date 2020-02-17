def print_response(response):
    if response is not None:
        print('status  :', response.status_code,
              '\nresponse:', response.json(), '\n')
    return response


def wrap_response(req_func):
    def wrap(*args, **kwargs):
        return print_response(req_func(*args, **kwargs))

    return wrap

