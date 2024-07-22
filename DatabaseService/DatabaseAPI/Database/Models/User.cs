﻿using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace DatabaseAPI.Database.Models;

public class User
{
	/*
    public User(string name, string surname, string login, string password, string roleKey)
    {
        Name = name;
		Surname = surname;
		Login = login;
		Password = password;
		RoleKey = roleKey;
    }
	*/

    [Key]
	[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
	public int Id { get; set; }

	public string Name { get; set; }

	public string Surname { get; set; }
	public int Balance { get; set; }

	public string Login { get; set; }

	public string Password { get; set; }

	public string RoleKey { get; set; }

	[ForeignKey(nameof(RoleKey))]
	public Role? Role { get; set; }
}