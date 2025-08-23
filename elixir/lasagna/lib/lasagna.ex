defmodule Lasagna do
  @expected_overtime 40
  @preparation_time 2
  # Please define the 'expected_minutes_in_oven/0' function
  def expected_minutes_in_oven do
    @expected_overtime
  end

  # Please define the 'remaining_minutes_in_oven/1' function
  def remaining_minutes_in_oven(time_in_oven) do
    @expected_overtime - time_in_oven
  end

  # Please define the 'preparation_time_in_minutes/1' function
  def preparation_time_in_minutes(layers) do
    layers * @preparation_time
  end

  # Please define the 'total_time_in_minutes/2' function
  def total_time_in_minutes(layer, time_in_oven) do
    time_in_oven + preparation_time_in_minutes(layer)
  end

  # Please define the 'alarm/0' function
  def alarm do
    "Ding!"
  end
end
