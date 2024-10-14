class Calculator
  def add(a, b)
    a + b
  end

  def subtract(a, b)
    a - b
  end

  def multiply(a, b)
    a * b
  end

  def divide(a, b)
    return 'Error: Division by zero' if b.zero?
    a.to_f / b
  end

  def exponentiate(base, exponent)
    base**exponent
  end

  def square_root(a)
    return 'Error: Negative number' if a.negative?
    Math.sqrt(a)
  end
end

calc = Calculator.new
puts calc.add(5, 3)
puts calc.subtract(5, 3)
puts calc.multiply(5, 3)
puts calc.divide(5, 0)
puts calc.divide(5, 2)
puts calc.exponentiate(2, 3)
puts calc.square_root(16)
puts calc.square_root(-4)
