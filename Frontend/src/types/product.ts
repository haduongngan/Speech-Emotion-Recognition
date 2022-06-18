export interface IProduct {
  color: number;
  errorDescription: string;
  id: string;
  image: string;
  name: string;
  sku: string;
}
export interface IProductColor {
  id: number;
  name: string;
}
export interface IProductLabelHeader {
  label: string;
}