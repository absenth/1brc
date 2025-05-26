#!/usr/bin/env python3
"""
* data: {city: {min: temp, max: temp, sum: temp, count: temp}}
* Read the file by line
* If city is new, add city to dictionary
"""
import pprint



def main(data_file):
    temps=get_temps(data_file)

    for city, data in sorted(temps.items()):
        print(f"{city}={data['min_temp']}/{data['sum_temp'] / data['temp_count']:.1f}/{data['max_temp']}")




def get_temps(data_file):
    main_dict={}

    with open(data_file, 'r') as file:
        for line in file:
            line = line.strip()
            if not line:
                continue  # Skip empty lines

            try:
                city_part, temp_part = line.split(';')
            except ValueError:
                continue  # Skip lines with invalid format

            city = city_part.strip()
            temp = float(temp_part.strip())

            if city not in main_dict:
                # Initialize the entry for a new city
                main_dict[city] = {
                    'min_temp': temp,
                    'max_temp': temp,
                    'sum_temp': temp,
                    'temp_count': 1
                }
            else:
                # Update existing city data
                existing = main_dict[city]
                existing_min = existing['min_temp']
                existing_max = existing['max_temp']

                # Update min and max
                new_min = min(existing_min, temp)
                new_max = max(existing_max, temp)

                # Update sum and count
                existing['sum_temp'] += temp
                existing['temp_count'] += 1

                # Set the updated min and max values
                existing['min_temp'] = new_min
                existing['max_temp'] = new_max

    return main_dict



if __name__ == "__main__":
    main("data/10ksample.txt")
