class ProductType:
  Win = "Win"
  Common = "Common"


def getProductType(hscode):
  if hscode == "11":
    return ProductType.Common
  else:
    return ProductType.Win

class TaxCompute(object):
  """
    ## 包含各种税费的计算方法
  """
  @classmethod
  def common(self, price, qty):
    radio = 0.2
    return price * qty * radio
  @classmethod
  def win(self, price, qty):
    radio = 0.05
    return price * qty * radio

# 计算税费的查表
TaxComputeFuncMap = {
  0: {
    ProductType.Common: TaxCompute.common,
    ProductType.Win: TaxCompute.win
  },
  1: {
    ProductType.Common: TaxCompute.common,
    ProductType.Win: TaxCompute.win
  }
}

def ComputeTaxPrice(withTax, productType, price, qty):
  """
    ## 计算税费
  """
  taxFunc = TaxComputeFuncMap.get(withTax, {}).get(productType, None)
  if taxFunc is None:
    raise Exception("Invalid tax type")
  return taxFunc(price, qty)

def test():
  withTax = 0
  price, qty = 2000, 2
  productType = getProductType("111")
  print(ComputeTaxPrice(withTax, productType, price, qty))

if __name__ == '__main__':
  test()