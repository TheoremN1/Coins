using System.ComponentModel.DataAnnotations;

namespace DatabaseAPI.Database.Models;

public class Status(string key, string name)
{
	[Key]
	public string Key { get; set; } = key;
	[Required]
	public string Name { get; set; } = name;
}