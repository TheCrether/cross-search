import { Result } from "../interfaces";

export function search(): Promise<Result[]> {
  return new Promise(() => {
    return [{ id: 1, name: "lel", icon: "" }]
  });
}

export default search;