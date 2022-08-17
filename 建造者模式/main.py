from Item import Meal
from food import VegBurger, Coke, ChickenBurger, Pepsi

class MealBuilder():
  def prepareVegMeal(self):
    meal = Meal()
    meal.addItem(VegBurger())
    meal.addItem(Coke())
    return meal
  def prepareNonVegMeal(self):
    meal = Meal()
    meal.addItem(ChickenBurger())
    meal.addItem(Pepsi())
    return meal

def test():
  mealBuilder = MealBuilder()
  vegMeal = mealBuilder.prepareVegMeal()
  nonVegMeal = mealBuilder.prepareNonVegMeal()
  vegMeal.showItems()
  print(f"Cost of vegMeal is {vegMeal.getCost()}")
  nonVegMeal.showItems()
  print(f"Cost of nonVegMeal is {nonVegMeal.getCost()}")

if __name__ == '__main__':
  test()