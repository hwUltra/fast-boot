import request from "@/utils/request";
import { AxiosPromise } from "axios";
import { OrderForm, OrderQuery, OrderVO } from "./types";

export function createOrder(queryParams?: OrderQuery): AxiosPromise<DeptVO[]> {
  return request({
    url: "/sys/dept/list",
    method: "get",
    params: queryParams,
  });
}
