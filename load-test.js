import { sleep } from "k6";
import http from "k6/http";
import {Counter} from 'k6/metrics'

let failedTestCases = new Counter('failedTestCases');

export let options = {
  duration: "5m",
  vus: 500,
  thresholds: {
        http_req_duration: ["p(90)<300"],
        http_req_waiting: ["p(90)<300"],
        failedTestCases: [{threshold: 'count===0'}]
    }
};

export default function() {
  let params = {
    headers: { 'Host' : `${__ENV.LOAD_TEST_HOST}` }
  }

  let res = http.get(`${__ENV.LOAD_TEST_URL}/health`, params);
  failedTestCases.add(res.status !== 200); 
  sleep(3);
}
