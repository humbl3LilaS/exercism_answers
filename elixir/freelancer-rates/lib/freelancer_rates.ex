defmodule FreelancerRates do
  @working_hours 8.0
  @working_days 22

  def daily_rate(hourly_rate) do
    hourly_rate * @working_hours
  end

  def apply_discount(before_discount, discount) do
    before_discount * (1 - conv_percentile(discount))
  end

  def monthly_rate(hourly_rate, discount) do
    (@working_days * daily_rate(hourly_rate)) |> apply_discount(discount) |> ceil()
  end

  def days_in_budget(budget, hourly_rate, discount) do
    discounted_rate = hourly_rate |> daily_rate() |> apply_discount(discount)
    (budget / discounted_rate) |> Float.floor(1)
  end

  defp conv_percentile(x) do
    x * 0.01
  end
end
