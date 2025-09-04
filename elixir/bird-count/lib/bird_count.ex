defmodule BirdCount do
  def today([head | _]) do
    head
  end

  def today([]) do
    nil
  end

  def increment_day_count([head | tail]) do
    [head + 1 | tail]
  end

  def increment_day_count([]) do
    [1]
  end

  def has_day_without_birds?([0 | _]) do
    true
  end

  def has_day_without_birds?([_ | tail]) do
    has_day_without_birds?(tail)
  end

  def has_day_without_birds?([]) do
    false
  end

  def total(list) do
    total(0, list)
  end

  def total(acc, [head | tail]) do
    total(acc + head, tail)
  end

  def total(acc, []) do
    acc
  end

  def busy_days(list) do
    busy_days(0, list)
  end

  def busy_days(acc, [head | tail]) when head >= 5 do
    busy_days(acc + 1, tail)
  end

  def busy_days(acc, [_ | tail]) do
    busy_days(acc, tail)
  end

  def busy_days(acc, []) do
    acc
  end
end
