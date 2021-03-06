package main

var listItemTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!-- Generated with glade 3.22.1 -->
<interface>
  <requires lib="gtk+" version="3.20"/>
  <object class="GtkBox" id="list_box">
    <property name="width_request">500</property>
    <property name="height_request">20</property>
    <property name="visible">True</property>
    <property name="can_focus">False</property>
    <property name="halign">start</property>
    <property name="valign">baseline</property>
    <child>
      <object class="GtkImage">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="stock">gtk-missing-image</property>
      </object>
      <packing>
        <property name="expand">False</property>
        <property name="fill">True</property>
        <property name="position">0</property>
      </packing>
    </child>
    <child>
      <object class="GtkLabel">
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="label" translatable="yes">label</property>
      </object>
      <packing>
        <property name="expand">False</property>
        <property name="fill">False</property>
        <property name="position">1</property>
      </packing>
    </child>
    <style>
      <class name="listBox"/>
    </style>
  </object>
</interface>
`
