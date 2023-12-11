import http from 'k6/http';
import { sleep } from 'k6';
export const options = {
  vus: 50,
  duration: '10s',
};
export default function () {
    http.get('http://localhost:1323/order/12cc7ad7-dd7a-4294-b2b2-34adbe209938');
    sleep(1);
}